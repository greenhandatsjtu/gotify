package database

import (
	"github.com/greenhandatsjtu/notifier/config"
	"github.com/greenhandatsjtu/notifier/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

var Db *gorm.DB
var err error

// 连接数据库
func ConnectDBandInit() {
	home, _ := homedir.Dir()
	Db, err = gorm.Open("sqlite3", filepath.Join(home, ".gotify", "gotify.db"))
	if err != nil {
		log.Fatal(err)
	}
	//迁移数据库
	Db.AutoMigrate(&models.Source{}, &models.Rss{}, &models.Json{}, &models.Plugin{})
	//需要初始化数据库
	if config.Config.Database.Init {
		//先drop现有的table
		Db.DropTableIfExists(&models.Source{}, &models.Rss{}, &models.Json{}).AutoMigrate(&models.Source{}, &models.Rss{}, &models.Json{}, &models.Plugin{})
		InitSources()
		InitRss()
		InitJson()
		log.Fatal("初始化数据库完成，请在配置文件中取消初始化设置") // 初始化完成后自动退出程序
	}
}
