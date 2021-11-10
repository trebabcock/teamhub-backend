package db

import (
	"log"
	"os"
	"teamhub-backend/app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func Init() *gorm.DB {
	dbURI := os.Getenv("DSN")
	//dbURI := "host=localhost port=5432 dbname=postgres user=postgres password=fglhub sslmode=disable"
	log.Println("Connecting to database...")
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Println("Could not connect to database")
		log.Fatal(err.Error())
	}
	log.Println("Connected to database")
	log.Println("Migrating database...")
	mdb := migrate(db)
	log.Println("Database Migrated")
	return mdb
}

func migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Settings{}, &model.Message{}, &model.Channel{}, &model.Post{}, &model.Comment{}, &model.Space{}, &model.SpaceItem{})
	return db
}
