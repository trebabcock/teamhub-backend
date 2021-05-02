package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Roles    []uint `json:"roles"`
	Banned   bool   `json:"banned"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	gorm.Model
	Name        string          `gorm:"unique" json:"name"`
	Color       string          `json:"color"`
	Permissions map[string]bool `json:"permissions"`
}

type UserSettings struct {
	gorm.Model
	User    uint     `json:"user"`
	Options Settings `json:"options"`
}

type Settings struct {
	Options map[string]bool `json:"options"`
}
