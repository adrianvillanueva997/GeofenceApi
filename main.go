package main

import (
	"log"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/config"
	"github.com/adrianvillanueva997/GeofenceApi/pkg/debug"
)

func main() {
	debug.InitDebugLogging()
	app := config.InitAPI()
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
