package models

import (
	"gorm.io/gorm"
	"github.com/megamsquare/dsanda/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `gorm:"not null" json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}