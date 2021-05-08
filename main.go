package main


import (
	app "fglhub-backend/app"
	db "fglhub-backend/db"
)

func main() {
	db := db.GetConfig()

	app := &app.App{}
	app.Init(db)
	app.Run(":2814")
}
