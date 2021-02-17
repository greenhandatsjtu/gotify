package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/greenhandatsjtu/notifier/database"
	"github.com/greenhandatsjtu/notifier/models"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

//获取所有爬虫源
func GetSources(c *gin.Context) {
	var sources []models.Source
	database.Db.Find(&sources)
	c.JSON(http.StatusOK, sources)
}

//添加爬虫源
func PostSource(c *gin.Context) {
	var source models.Source
	var newSource models.NewSource
	if err := c.Bind(&newSource); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	blockPattern := regexp.MustCompile(`:nth-child\(\d\)`)
	source.Block = blockPattern.ReplaceAllString(newSource.Block, "")

	baseUrlPattern := regexp.MustCompile(`(?U)^https?://.+/`)
	source.BaseUrl = baseUrlPattern.FindString(newSource.Url)

	hrefPattern := regexp.MustCompile(`^<.*href.*>.*>$`)
	if hrefPattern.MatchString(newSource.Html) {
		source.Href = "href"
	} else {
		log.Error("请正确填写各项内容！")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	titlePattern := regexp.MustCompile(`^<.*title.*>.*>$`)
	if titlePattern.MatchString(newSource.Html) {
		source.Title = "title"
	} else {
		titlePattern = regexp.MustCompile(`(?U)^<.*>.*\S+.*<`)
		if titlePattern.MatchString(newSource.Html) {
			source.Title = "Text()"
		} else {
			log.Error("请正确填写各项内容！")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
	source.Name = newSource.Name
	source.Url = newSource.Url
	source.Enabled = &newSource.Enabled
	source.Frequency = newSource.Frequency
	database.Db.Create(&source)
}

//更新爬虫源
func UpdateSource(c *gin.Context) {
	id := c.Params.ByName("id") //根据id查找
	var source, newSource models.Source
	database.Db.Where(id).First(&source)
	if err := c.Bind(&newSource); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//更新
	database.Db.Model(&source).Updates(newSource)
}

//删除爬虫源
func DeleteSource(c *gin.Context) {
	id := c.Params.ByName("id") //根据id查找
	var source models.Source
	if err := database.Db.Where(id).First(&source).Error; err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	database.Db.Delete(&source)
}

//禁用所有爬虫源
func DisableSource(c *gin.Context) {
	database.Db.Model(&models.Source{}).Updates(map[string]interface{}{"enabled": false})
	c.JSON(http.StatusOK, gin.H{"msg": "成功禁用所有爬虫源"})
}

//获取所有rss源
func GetRssSources(c *gin.Context) {
	var sources []models.Rss
	database.Db.Find(&sources)
	c.JSON(http.StatusOK, sources)
}

//添加rss源
func PostRssSource(c *gin.Context) {
	var source models.Rss
	if err := c.Bind(&source); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	database.Db.Create(&source)
}

//更新rss源
func UpdateRssSource(c *gin.Context) {
	id := c.Params.ByName("id") //根据id查找
	var source, newSource models.Rss
	database.Db.Where(id).First(&source)
	if err := c.Bind(&newSource); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	database.Db.Model(&source).Update(newSource)
}

//删除rss源
func DeleteRssSource(c *gin.Context) {
	id := c.Params.ByName("id")
	var source models.Rss
	if err := database.Db.Where(id).First(&source).Error; err != nil {
		log.Warn(err)
		c.AbortWithStatus(404)
		return
	}
	database.Db.Delete(&source)
}

//禁用所有rss源
func DisableRssSource(c *gin.Context) {
	database.Db.Model(&models.Rss{}).Updates(map[string]interface{}{"enabled": false})
	c.JSON(http.StatusOK, gin.H{"msg": "成功禁用所有RSS源"})
}

//获取所有json源
func GetJsonSources(c *gin.Context) {
	var sources []models.Json
	database.Db.Find(&sources)
	c.JSON(http.StatusOK, sources)
}

//添加json源
func PostJsonSource(c *gin.Context) {
	var source models.Json
	if err := c.Bind(&source); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	database.Db.Create(&source)
}

//更新json源
func UpdateJsonSource(c *gin.Context) {
	id := c.Params.ByName("id")
	var source, newSource models.Json
	database.Db.Where(id).First(&source)
	if err := c.Bind(&newSource); err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	database.Db.Model(&source).Update(newSource)
}

//删除json源
func DeleteJsonSource(c *gin.Context) {
	id := c.Params.ByName("id")
	var source models.Json
	if err := database.Db.Where(id).First(&source).Error; err != nil {
		log.Warn(err)
		c.AbortWithStatus(404)
		return
	}
	database.Db.Delete(&source)
}

//禁用所有json源
func DisableJsonSource(c *gin.Context) {
	database.Db.Model(&models.Json{}).Updates(map[string]interface{}{"enabled": false})
	c.JSON(http.StatusOK, gin.H{"msg": "成功禁用所有json源"})
}
