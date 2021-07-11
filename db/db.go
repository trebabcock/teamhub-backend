package db

import (
	"fmt"
	"log"
	"os"
	"teamhub-backend/app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func Init() *gorm.DB {
	dbURI := os.Getenv("DB_URI")
	fmt.Println("Connecting to database...")
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Println("Could not connect to database")
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to database")
	fmt.Println("Migrating database...")
	mdb := migrate(db)
	fmt.Println("Database Migrated")
	return mdb
}

func migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Settings{}, &model.Message{}, &model.Channel{}, &model.Post{}, &model.Comment{}, &model.Space{}, &model.SpaceItem{})
	return db
}
