package models

import "github.com/jinzhu/gorm"

type Plugin struct {
	gorm.Model
	Name      string `gorm:"not_null;default:plugin" json:"name" form:"name"`
	Exec      string `gorm:"not_null" json:"exec" form:"exec"`
	Frequency int    `gorm:"not_null;default:30" json:"frequency" form:"frequency"`
	Enabled   *bool  `gorm:"not_null;default:0" json:"enabled" form:"enabled"`
}
