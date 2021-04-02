package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPoints(ctx *fiber.Ctx) error {
	response := models.PolygonResponse{
		Status:  fiber.StatusOK,
		Polygon: polygon.Geometry,
	}
	return ctx.Status(response.Status).JSON(response)
}
