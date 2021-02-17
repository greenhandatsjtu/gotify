package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greenhandatsjtu/gotify/config"
	"github.com/greenhandatsjtu/gotify/database"
	_ "github.com/greenhandatsjtu/gotify/logs"
	"github.com/greenhandatsjtu/gotify/models"
	"github.com/greenhandatsjtu/gotify/routes"
	"github.com/greenhandatsjtu/gotify/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
)

func main() {
	//加载配置
	config.LoadConfig()

	//连接并初始化数据库
	database.ConnectDBandInit()

	//make退出和重启信号的channel
	quit := make(chan bool)
	defer close(quit)
	reboot := make(chan bool)
	defer close(reboot)
	//make一个缓冲为64的消息信道
	msgs := make(chan models.NewMsg, 64)
	defer close(msgs)

	//如果配置中启用了web界面
	if config.Config.Server.Enable {
		gin.SetMode(gin.ReleaseMode)
		r := routes.InitRoutes(quit, reboot) // 初始化路由
		go r.Run(config.Config.Server.IP + ":" + strconv.Itoa(config.Config.Server.Port))
		//在浏览器中打开web界面
		log.Infoln("web管理页面运行在:", config.Config.Server.Port, "端口")
		//exec.Command("chromium", config.Config.Server.IP+":"+strconv.Itoa(config.Config.Server.Port)).Start()
	}

	//创建http client
	utils.NewHttpClient(config.Config.HTTPClient.Timeout)

	//启用通知的goroutine
	home, _ := homedir.Dir()
	go utils.Notify(msgs, filepath.Join(home, "images", "logo.png"))

	loggedIn := false
	needQuit := false
	for {
		needQuit = startSync(quit, reboot, msgs, &loggedIn)
		if needQuit {
			break
		}
	}
}
