package model

import "github.com/jinzhu/gorm"

type Space struct {
	gorm.Model
	Owner   uint   `json:"owner"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
	Content []uint `json:"content"`
}

type SpaceItem struct {
	gorm.Model
	Type    string `json:"type"`
	Content string `json:"content"`
	Space   uint   `json:"space"`
}
