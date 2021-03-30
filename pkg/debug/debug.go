package debug

import (
	"log"
	"os"
)

const DEBUGFILE = "logs/debug.log"

func InitDebugLogging() {
	logFile, err := os.OpenFile(DEBUGFILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
