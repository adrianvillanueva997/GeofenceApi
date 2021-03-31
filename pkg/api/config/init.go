package config

import (
	v12 "github.com/adrianvillanueva997/GeofenceApi/pkg/api/endpoints/v1"
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/middleware"
	"github.com/gofiber/fiber/v2"
)

// Initialise API server
func InitAPI() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	setV1APIRoutes(api)
	middleware.InitAPILogging(app)
	middleware.InitAPILimiter(app)
	return app
}

// Set API routes
func setV1APIRoutes(api fiber.Router) {
	v1 := api.Group("/v1")
	v1.Get("/coordinates", v12.GetPoints)
	v1.Post("/updateCoordinate", v12.UpdatePoint)
	v1.Post("/inside", v12.CheckDeviceInsidePoint)
	v1.Get("/health", v12.Health)
}
