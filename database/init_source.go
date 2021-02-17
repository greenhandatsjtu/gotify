//用于初始化各源的数据库
package database

import (
	"github.com/greenhandatsjtu/gotify/models"
)

//初始化订阅源（爬虫）
func InitSources() {
	var source models.Source

	source = models.Source{
		Url:       "https://neets.cc/articles/2",
		BaseUrl:   "",
		Block:     "h1 a[href]",
		Title:     "Text()",
		Href:      "href",
		Frequency: 100,
		Mark:      "https://mp.weixin.qq.com/s/DGOxFCFYWzereez_SMoW4A",
	}
	Db.Create(&source)
	Db.Model(&source).Updates(map[string]interface{}{"enabled": true})

	source = models.Source{
		Url:       "http://xsb.seiee.sjtu.edu.cn/xsb/list/2496-1-20.htm",
		BaseUrl:   "http://xsb.seiee.sjtu.edu.cn",
		Block:     "a[title]",
		Title:     "title",
		Href:      "href",
		Frequency: 120,
		Mark:      "/xsb/detail/2496_4116.htm?nocache=1",
	}
	Db.Create(&source)
	Db.Model(&source).Updates(map[string]interface{}{"enabled": true})

	source = models.Source{
		Url:       "https://movie.douban.com/review/best/",
		BaseUrl:   "",
		Block:     "h2 > a[href]",
		Title:     "Text()",
		Href:      "href",
		Frequency: 140,
		Mark:      "https://movie.douban.com/review/12435989/",
	}
	Db.Create(&source)
	Db.Model(&source).Updates(map[string]interface{}{"enabled": true})

	source = models.Source{
		Url:       "http://zhuixinfan.com/main.php",
		BaseUrl:   "http://zhuixinfan.com/",
		Block:     ".f1 > a[target]",
		Title:     "Text()",
		Href:      "href",
		Frequency: 90,
		Mark:      "main.php?mod=viewresource&sid=11223",
	}
	Db.Create(&source)
	Db.Model(&source).Updates(map[string]interface{}{"enabled": true})

	source = models.Source{
		Url:       "https://www.v2ex.com/",
		BaseUrl:   "https://www.v2ex.com",
		Block:     ".item_title > a",
		Title:     "Text()",
		Href:      "href",
		Frequency: 12,
		Mark:      "/t/656566#reply111",
	}
	Db.Create(&source)
	Db.Model(&source).Updates(map[string]interface{}{"enabled": true})

	source = models.Source{
		Url:       "https://www.solidot.org/",
		BaseUrl:   "https://www.solidot.org",
		Block:     ".bg_htit > h2> a",
		Title:     "Text()",
		Href:      "href",
		Frequency: 31,
		Mark:      "/story?sid=63941",
	}
	Db.Create(&source)
	Db.Model(&source).Updates(map[string]interface{}{"enabled": true})
}

//初始化RSS feed
func InitRss() {
	var rss models.Rss
	rss = models.Rss{
		Url:  "https://www.cnbeta.com/backend.php",
		Mark: "https://www.cnbeta.com/articles/tech/960773.htm",
	}
	Db.Create(&rss)
	Db.Model(&rss).Updates(map[string]interface{}{"enabled": true})

	rss = models.Rss{
		Url:  "https://blog.golang.org/feed.atom",
		Mark: "tag:blog.golang.org,2013:blog.golang.org/pandemic",
	}
	Db.Create(&rss)
	Db.Model(&rss).Updates(map[string]interface{}{"enabled": true})

	rss = models.Rss{
		Url:  "https://pt.sjtu.edu.cn/torrentrss.php?rows=10&cat413=1&cat431=1",
		Mark: "06297f0b2cedeb2e31b4dd681c1f8cccae586039",
	}
	Db.Create(&rss)
	Db.Model(&rss).Updates(map[string]interface{}{"enabled": true})

	rss = models.Rss{
		Url:  "https://share.dmhy.org/topics/rss/rss.xml",
		Mark: "http://share.dmhy.org/topics/view/537863_LoliHouse_Koisuru_Asteroid_-_11_WebRip_1080p_HEVC-10bit_AAC.html",
	}
	Db.Create(&rss)
	Db.Model(&rss).Updates(map[string]interface{}{"enabled": true})
}

//初始化json
func InitJson() {
	var jsonSource models.Json
	jsonSource = models.Json{
		Url: "https://ncov-rss.qgis.me/api/messages?limit=10",
		//BaseUrl:   "",
		//Icon:      "ncov.jpg",
		Block: "messages",
		//Title:     "",
		//Message:   "message",
		//Href:      "",
		Mark:      "id",
		Frequency: 30,
		NewMark:   "4119",
	}
	Db.Create(&jsonSource)
	Db.Model(&jsonSource).Updates(map[string]interface{}{
		"enabled": true,
		"message": "message",
	})
}
