package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/greenhandatsjtu/notifier/database"
	"github.com/greenhandatsjtu/notifier/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//获取所有插件
func GetPlugins(c *gin.Context) {
	var plugins []models.Plugin
	database.Db.Find(&plugins)
	c.JSON(http.StatusOK, plugins)
}

//添加插件
func PostPlugin(c *gin.Context) {
	var plugin models.Plugin
	if err := c.Bind(&plugin); err != nil {
		log.Warn(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}
	database.Db.Create(&plugin)
}

//更新插件
func UpdatePlugin(c *gin.Context) {
	id := c.Params.ByName("id")
	var plugin, newPlugin models.Plugin
	database.Db.Where(id).First(&plugin)
	if err := c.Bind(&newPlugin); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	database.Db.Model(&plugin).Update(newPlugin)
}

//删除插件
func DeletePlugin(c *gin.Context) {
	id := c.Params.ByName("id")
	var plugin models.Plugin
	if err := database.Db.Where(id).First(&plugin).Error; err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	database.Db.Delete(&plugin)
	c.JSON(http.StatusOK, gin.H{"msg": "成功移除插件"})
}

//禁用所有插件
func DisablePlugins(c *gin.Context) {
	database.Db.Model(&models.Plugin{}).Updates(map[string]interface{}{"enabled": false})
	c.JSON(http.StatusOK, gin.H{"msg": "成功禁用所有插件"})
}
