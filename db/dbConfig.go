package db

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Dialect		string
	Host		string
	Port		string
	Username	string
	Name		string
	Password	string
	SSLMode		string
}

func GetConfig() *Config {
	config := Config{
		Dialect: os.Getenv("DB_DIALECT"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Name: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	return &config
}
