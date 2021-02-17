package models

import "github.com/jinzhu/gorm"

type Json struct {
	gorm.Model
	Name      string  `gorm:"not_null;default:json" json:"name" form:"name"`
	Url       string  `gorm:"size:1000;not_null" json:"url" form:"url"`
	BaseUrl   *string `gorm:"size:1000" json:"base_url" form:"base_url"`
	Block     string  `gorm:"not_null" json:"block" form:"block"`
	Title     *string `json:"title" form:"title"`
	Message   *string `json:"message" form:"message"`
	Href      *string `json:"href" form:"href"`
	Mark      string  `json:"mark" form:"mark"`
	NewMark   string  `json:"new_mark" form:"new_mark"`
	Frequency int     `gorm:"not_null;default:30" json:"frequency" form:"frequency"`
	Enabled   *bool   `gorm:"not_null;default:0" json:"enabled" form:"enabled"`
}
