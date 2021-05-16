package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
	Banned   bool   `json:"banned"`
}

type PublicUser struct {
	Name     string
	Username string
	Roles    []Role
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	gorm.Model
	Name        string       `gorm:"unique" json:"name"`
	Color       string       `json:"color"`
	Permissions []Permission `gorm:"many2many:user_permissions;" json:"permissions"`
}

type Permission struct {
	Option  string `json:"option"`
	Enabled bool   `json:"enabled"`
}

type UserSettings struct {
	gorm.Model
	User    uint       `json:"user"`
	Options []Settings `gorm:"many2many:user_settings;" json:"options"`
}

type Settings struct {
	Option  string `json:"options"`
	Enabled bool   `json:"enabled"`
}
