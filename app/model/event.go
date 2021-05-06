package model

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	User    uint   `json:"user"`
	Type    string `json:"type"`
	Message string `josn:"message"`
}
