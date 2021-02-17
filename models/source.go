package models

import (
	"github.com/jinzhu/gorm"
)

type Source struct {
	gorm.Model
	Name      string `gorm:"not_null;default:source" json:"name" form:"name"`
	Url       string `gorm:"size:1000;not_null" json:"url" form:"url"`
	BaseUrl   string `gorm:"size:1000" json:"base_url" form:"base_url"`
	Block     string `gorm:"not_null" json:"block" form:"block"`
	Title     string `gorm:"not_null" json:"title" form:"title"`
	Href      string `gorm:"size:1000;not_null" json:"href" form:"href"`
	Frequency int    `gorm:"not_null;default:30" json:"frequency" form:"frequency"`
	Mark      string `gorm:"size:1000" json:"mark" form:"mark"`
	Enabled   *bool  `gorm:"not_null;default:0" json:"enabled" form:"enabled"`
}

//新建爬虫源所用struct
type NewSource struct {
	Name      string `json:"name" form:"name"`
	Url       string `json:"url" form:"url"`
	Html      string `json:"html" form:"html"`
	Block     string `json:"block" form:"block"`
	Frequency int    `json:"frequency" form:"frequency"`
	Enabled   bool   `json:"enabled" form:"enabled"`
}
