package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const ReFile = "logs/requests.log"

var loggerConfig = logger.Config{
	Next:         nil, // Function to skip middleware when return True
	Format:       "[${time}] ${ip} - ${status} - ${latency} - ${method} - ${path}\n",
	TimeFormat:   "15:04:05",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond, // Interval between log writes
	Output:       writeToLogs(),
}

func writeToLogs() *os.File {
	file, err := os.OpenFile(ReFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}

//
func InitAPILogging(app *fiber.App) {
	app.Use(logger.New(loggerConfig))
}
