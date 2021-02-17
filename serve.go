package main

import (
	"fmt"
	"github.com/greenhandatsjtu/notifier/config"
	"github.com/greenhandatsjtu/notifier/database"
	"github.com/greenhandatsjtu/notifier/models"
	"github.com/greenhandatsjtu/notifier/sync"
	"github.com/greenhandatsjtu/notifier/utils"
	log "github.com/sirupsen/logrus"
)

//拉去消息的主要函数，返回true是表示退出程序，返回false表示重启
func startSync(quit, reboot chan bool, msgs chan models.NewMsg, loggedIn *bool) bool {
	//开启go程读取键盘输入的命令
	rebootFlag := false
	go func() {
		var input string
		for {
			fmt.Scan(&input)
			if input == "quit" {
				quit <- true
				return
			} else if input == "reboot" {
				reboot <- true
				rebootFlag = true
				return
			}
		}
	}()

	//如果rss源启用
	if config.Config.Rss {
		var rssSources []models.Rss
		//从数据库读取rss源
		if err := database.Db.Find(&rssSources).Error; err != nil {
			log.Fatal(err)
		}
		for _, source := range rssSources {
			//同样犯了之前的错误导致4个go程的source一样，如果直接使用go func(){}，则source其实是同一个地址，应该要将source传入函数
			if *source.Enabled {
				//为每一个启用的rss源开启go程独立拉取信息
				go sync.SyncRss(source, msgs)
			}
		}
	}

	//自定义爬虫消息源，从数据库查询
	if config.Config.Crawler {
		var sourses []models.Source
		if err := database.Db.Find(&sourses).Error; err != nil {
			log.Error(err)
		}
		for _, source := range sourses {
			if *source.Enabled {
				//为每一个启用的爬虫源开启go程独立拉取信息
				go sync.LoopSyncSource(source, msgs)
			}
		}
	}

	//自定义插件，从数据库查询
	if config.Config.Plugin {
		var plugins []models.Plugin
		if err := database.Db.Find(&plugins).Error; err != nil {
			log.Error(err)
		}
		for _, plugin := range plugins {
			if *plugin.Enabled {
				//为每一个启用的插件开启go程
				go sync.LoopSyncPlugin(plugin, msgs)
			}
		}
	}

	//自定义JSON消息源，从数据库查询
	if config.Config.JSON {
		var jsonSourses []models.Json
		if err := database.Db.Find(&jsonSourses).Error; err != nil {
			log.Error(err)
		}

		for _, source := range jsonSourses {
			if *source.Enabled {
				//为每一个启用的json源开启go程独立拉取信息
				go sync.LoopGetJson(msgs, utils.Client, source)
			}
		}
	}

	//读取没有阻塞channel
	select {
	case <-reboot:
		log.Info("重启服务中...")
		return false
	case <-quit:
		log.Info("退出程序...")
		return true
	}
}
