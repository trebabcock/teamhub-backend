package model

import (
	"github.com/jinzhu/gorm"
)

type Packet struct {
	PacketType        string `json:"type"`
	PacketSender      string `json:"sender"`
	PacketDestination string `json:"destination"`
}

type Message struct {
	gorm.Model
	AuthorID     string `json:"author_id"`
	UUID         string `json:"uuid"`
	Time         string `json:"time"`
	DesinationID string `json:"destination_id"`
	Private      bool   `json:"private"`
	Type         string `json:"type"`
	Content      string `json:"content"`
}

type Channel struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}
