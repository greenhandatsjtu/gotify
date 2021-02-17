package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/greenhandatsjtu/notifier/controllers"
	_ "github.com/greenhandatsjtu/notifier/statik"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// 初始化gin的路由
func InitRoutes(quit, reboot chan bool) *gin.Engine {
	//m := melody.New() //websocket库，已弃用
	r := gin.Default()
	//加载HTML文件
	//r.LoadHTMLFiles(config.Config.Root + config.Config.Server.Static + config.Config.Server.Index)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	r.StaticFS("/index", statikFS)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/index")
	})

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "OK"})
	})

	r.GET("/reboot", func(context *gin.Context) {
		//context.JSON(200, gin.H{"status": "OK"})
		//发出重启信号
		reboot <- true
	})

	r.GET("/quit", func(context *gin.Context) {
		//context.JSON(200, gin.H{"status": "OK"})
		//发出退出信号
		quit <- true
	})

	//查询路由
	r.GET("/sources", controllers.GetSources)
	r.GET("/rssSources", controllers.GetRssSources)
	r.GET("/jsonSources", controllers.GetJsonSources)
	r.GET("/plugins", controllers.GetPlugins)

	//添加路由
	r.POST("/source", controllers.PostSource)
	r.POST("/rssSource", controllers.PostRssSource)
	r.POST("/jsonSource", controllers.PostJsonSource)
	r.POST("/plugin", controllers.PostPlugin)

	//更新路由
	r.POST("/source/:id", controllers.UpdateSource)
	r.POST("/rssSource/:id", controllers.UpdateRssSource)
	r.POST("/jsonSource/:id", controllers.UpdateJsonSource)
	r.POST("/plugin/:id", controllers.UpdatePlugin)

	//删除路由
	r.GET("/source/:id", controllers.DeleteSource)
	r.GET("/rssSource/:id", controllers.DeleteRssSource)
	r.GET("/jsonSource/:id", controllers.DeleteJsonSource)
	r.GET("/plugin/:id", controllers.DeletePlugin)

	//禁用路由
	r.GET("/sourceDisable", controllers.DisableSource)
	r.GET("/rssSourceDisable", controllers.DisableRssSource)
	r.GET("/jsonSourceDisable", controllers.DisableJsonSource)
	r.GET("/pluginDisable", controllers.DisablePlugins)

	return r
}
