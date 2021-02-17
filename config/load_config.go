package config

import (
	"github.com/greenhandatsjtu/gotify/models"
	"github.com/greenhandatsjtu/gotify/utils"
	"github.com/jinzhu/configor"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var Config models.Config

//加载配置文件
func LoadConfig() {
	home, _ := homedir.Dir()
	dir := filepath.Join(home, ".gotify")
	if _, err := os.Stat(dir); err != nil {
		utils.InitApp()
	}
	if err := configor.Load(&Config, filepath.Join(dir, "config.json")); err != nil {
		log.Fatal("load config error: ", err)
	}
}
