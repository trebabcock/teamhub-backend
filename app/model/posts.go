package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Author   string    `json:"author"`
	Type     string    `json:"type"`
	Time     string    `json:"time"`
	Content  string    `json:"content"`
	UUID     string    `json:"id"`
	Comments []Comment `gorm:"many2many:post_comments;" json:"comments"`
}

type NPost struct {
	Author  string `json:"author"`
	Type    string `json:"type"`
	Time    string `json:"time"`
	Content string `json:"content"`
	UUID    string `json:"id"`
}

type Comment struct {
	gorm.Model
	Author  uint   `json:"author"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Post    uint   `json:"comments"`
}
