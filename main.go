package main

import (
	"log"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/config"
)

func main() {
	app := config.InitAPI()
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
