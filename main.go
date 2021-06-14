package main

import (
	app "teamhub-backend/app"
	db "teamhub-backend/db"
)

func main() {
	db := db.GetConfig()

	app := &app.App{}
	app.Init(db)
	app.Run(":2814")
}
