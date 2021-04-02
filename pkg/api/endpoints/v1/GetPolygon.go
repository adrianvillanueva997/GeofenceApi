package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPolygon(ctx *fiber.Ctx) error {
	response := models.PolygonResponse{
		Status:  fiber.StatusOK,
		Polygon: polygon.Geometry,
		Message: "OK",
	}
	if polygon.Geometry.Coordinates == nil {
		response.Status = fiber.StatusNotFound
		response.Message = "Polygon not loaded yet."
	}
	return ctx.Status(response.Status).JSON(response)
}
