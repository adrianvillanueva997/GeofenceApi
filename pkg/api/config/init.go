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
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/coordinates", endpoints.GetPoints)
	v1.Post("/updateCoordinate", endpoints.UpdatePoint)
	v1.Post("/inside", endpoints.CheckDeviceInsidePoint)
	v1.Get("/health", endpoints.Health)
}
