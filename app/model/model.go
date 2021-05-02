package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Role{}, &Settings{}, &Message{}, &Channel{}, &Post{}, &Comment{}, &Space{}, &SpaceItem{})
	return db
}
