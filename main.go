package main

import (
	app "teamhub-backend/app"
)

func main() {

	app := &app.App{}
	app.Init()
	app.Run(":2814")
}
