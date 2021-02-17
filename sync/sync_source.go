package sync

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/greenhandatsjtu/gotify/database"
	"github.com/greenhandatsjtu/gotify/models"
	"github.com/greenhandatsjtu/gotify/utils"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

//SyncSource 用来拉去订阅源（爬虫）
func syncSource(source models.Source, msgs chan models.NewMsg) {
	req, err := http.NewRequest("GET", source.Url, nil)
	if err != nil {
		log.Warn(err)
		return
	}
	//添加User-Agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10")
	resp, err := utils.Client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(source.Url + " 得到响应，开始解析...")
	defer resp.Body.Close()

	//解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
	}

	//遍历查找标题
	feeds := doc.Find(source.Block)
	var newMark string
	feeds.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		msg := models.NewMsg{}
		if source.Title == "Text()" {
			msg.Title = selection.Text()
		} else {
			var exists bool
			msg.Title, exists = selection.Attr(source.Title)
			if !exists {
				log.Warn("Title not fount")
				return true
			}
		}
		//先去掉标题的左右两边空格
		msg.Title = source.Name + "|" + utils.CompressTitle(msg.Title)
		//达到之前已获取的新闻时，停止
		var exists bool
		//获取原文地址
		msg.Message, exists = selection.Attr(source.Href)
		if !exists {
			log.Warn("Href not found")
			return true
		}
		if msg.Message == source.Mark {
			//             fmt.Print(source.Url," ", i,"\n")
			return false
		}
		//缓存标记
		if i == 0 {
			newMark = msg.Message
		}
		msg.Message = "<a href=\"" + source.BaseUrl + msg.Message + "\">原文链接</a>"
		msgs <- msg

		//不要接受太多的消息
		if i == 2 {
			return false
		}
		return true
	})
	//更新标记
	database.Db.Model(&source).Update(models.Source{Mark: newMark})
	//utils.Db.Save(&source)
}

//这里的source必须采用值传递而不能用指针，因为source的指针会随着查询而改变，最后会指向最后一向，导致
//所有source都是相同的
func LoopSyncSource(source models.Source, msgs chan models.NewMsg) {
	timer := make(chan bool)
	defer close(timer)
	//设置定时器
	go utils.TickTock(timer, time.Duration(source.Frequency)*time.Minute)
	for {
		<-timer
		syncSource(source, msgs)
	}
}

//SyncRss 用来更新RSS
func SyncRss(source models.Rss, msgs chan models.NewMsg) {
	log.Info("Getting feed " + source.Url)
	feed, err := rss.Fetch(source.Url)
	for {
		if err != nil {
			log.Error(err)
			break
		}
		for i, item := range feed.Items {
			//获取3条或到达上次看过的条目时停止
			if i == 2 || item.ID == source.Mark {
				break
			}
			//更新标记
			if i == 0 {
				database.Db.Model(&source).Update(models.Rss{Mark: item.ID})
				database.Db.Save(&source)
			}
			//移进消息队列
			message := item.Summary
			if message == "" {
				message = item.Content
			}
			if item.Link != "" {
				message = "<a href=\"" + item.Link + "\">原文链接</a>\n" + message
			}
			msgs <- models.NewMsg{
				Title:   source.Name + "|" + item.Title,
				Message: message,
				//Icon:    feed.Image.URL,
			}
		}

		//休眠一定时间后更新RSS
		sleepTime := feed.Refresh.Local().Sub(time.Now().Local())
		time.Sleep(sleepTime)

		log.Info("Updating feed " + source.Url)
		if err = feed.Update(); err != nil {
			log.Error(err)
		}
	}
}

//拉取json源
func getJson(client http.Client, url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10")
	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err)
		return nil
	}
	defer resp.Body.Close()
	json, err := ioutil.ReadAll(resp.Body)
	log.Info(url, " ", resp.StatusCode)
	return json
}

//每当计时器发出信号就循环拉取帖子
func LoopGetJson(msgs chan models.NewMsg, client http.Client, source models.Json) {
	timer := make(chan bool)
	defer close(timer)
	//设置定时器
	go utils.TickTock(timer, time.Duration(source.Frequency)*time.Minute)

	for {
		<-timer
		log.Info("拉取 " + source.Url + " ...")
		unread := false // 未读标志
		index := 0
		json := getJson(client, source.Url)
		if gjson.Valid(string(json)) {
			gjson.GetBytes(json, source.Block).ForEach(func(key, value gjson.Result) bool {
				//最多读取3条后停止
				if index == 2 {
					return false
				}
				index++
				if value.Get(source.Mark).String() != source.NewMark {
					unread = true
					message := new(string)
					href := new(string)
					//查找message
					if source.Message != nil {
						*message = value.Get(*source.Message).String()
					}
					//查找超链接
					if source.Href != nil {
						if source.BaseUrl != nil {
							*href = "<a href=\"" + *source.BaseUrl + value.Get(*source.Href).String() + "\">原文链接</a>"
						} else {
							*href = "<a href=\"" + value.Get(*source.Href).String() + "\">原文链接</a>"
						}
					}
					//查找标题
					if source.Title != nil {
						msgs <- models.NewMsg{
							Title:   source.Name + "|" + value.Get(*source.Title).String(),
							Message: *href + "\n" + *message,
						}
					} else {
						//存入消息队列
						msgs <- models.NewMsg{
							Title:   source.Name,
							Message: *href + "\n" + *message,
						}
					}
				} else {
					return false
				}
				return true
			})
			//更新标记
			database.Db.Model(&source).Update(models.Json{NewMark: gjson.GetBytes(json, source.Block+".0."+source.Mark).String()})
			if !unread {
				log.Info(source.Url + ": 无最新消息")
			}
		} else {
			log.Warn(source.Url + " invalid json.")
		}
	}
}
