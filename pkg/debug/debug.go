package debug

import (
	"log"
	"os"
)

const Defile = "logs/debug.log"

func InitDebugLogging() {
	logFile, err := os.OpenFile(Defile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
