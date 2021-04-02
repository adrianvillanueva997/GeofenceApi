package config

import (
	endpointsV1 "github.com/adrianvillanueva997/GeofenceApi/pkg/api/endpoints/v1"
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
	v1.Get("/polygon", endpointsV1.GetPolygon)
	v1.Put("/update", endpointsV1.UpdatePolygon)
	v1.Get("/devicestatus", endpointsV1.DeviceStatus)
	v1.Get("/health", endpointsV1.Health)
}
