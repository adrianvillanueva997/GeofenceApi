package main

import (
	"log"
	"os"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/config"
)

const DEBUGFILE = "logs/debug.log"

func initDebugLogging() {
	logFile, err := os.OpenFile(DEBUGFILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	initDebugLogging()
	app := config.InitAPI()
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
