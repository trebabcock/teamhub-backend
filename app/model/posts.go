package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Author   uint      `json:"author"`
	Type     string    `json:"type"`
	Content  string    `json:"content"`
	Comments []Comment `gorm:"many2many:post_comments;" json:"comments"`
}

type Comment struct {
	gorm.Model
	Author  uint   `json:"author"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Post    uint   `json:"comments"`
}
