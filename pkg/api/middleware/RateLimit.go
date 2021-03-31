package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var limiterConfig = limiter.Config{
	Next: nil,
	Max:  5, // Max number of connections
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
	Expiration: 1 * time.Minute,
	LimitReached: func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusTooManyRequests)
	},
}

func InitAPILimiter(app *fiber.App) {
	app.Use(limiter.New(limiterConfig))
}
