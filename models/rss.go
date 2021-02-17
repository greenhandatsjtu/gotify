package models

import "github.com/jinzhu/gorm"

type Rss struct {
	gorm.Model
	Name    string `gorm:"not_null;default:rss" json:"name" form:"name"`
	Url     string `gorm:"size:1000;not_null" json:"url" form:"url"`
	Mark    string `gorm:"size:1000" json:"mark" form:"mark"`
	Enabled *bool  `gorm:"not_null;default:0" json:"enabled" form:"enabled"`
}
