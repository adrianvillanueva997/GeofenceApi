package config

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/endpoints"
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/middleware"
	"github.com/gofiber/fiber/v2"
)

// Initialise API server
func InitAPI() *fiber.App {
	app := fiber.New()
	setRoutes(app)
	middleware.InitAPILogging(app)
	middleware.InitAPILimiter(app)
	return app
}

// Set API routes
func setRoutes(app *fiber.App) {
	api := app.Group("/v1")
	api.Get("/coordinates", endpoints.GetPoints)
	api.Post("/updateCoordinate", endpoints.UpdatePoint)
	api.Post("/inside", endpoints.CheckDeviceInsidePoint)
}
