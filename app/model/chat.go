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
	Author     string `json:"author"`
	UUID       string `json:"uuid"`
	Time       string `json:"time"`
	Desination string `json:"destination"`
	Type       string `json:"type"`
	Content    string `json:"content"`
}

type Channel struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
