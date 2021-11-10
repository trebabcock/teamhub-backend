package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	UUID     string `json:"id"`
}

type RUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PublicUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

type LoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UUID     string `json:"id"`
	Token    []byte `json:"access_token"`
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
