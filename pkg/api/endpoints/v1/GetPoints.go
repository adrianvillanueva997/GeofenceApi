package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPoints(ctx *fiber.Ctx) error {
	response := models.GetCoordinates{
		Status:              fiber.StatusOK,
		NumberOfCoordinates: len(coordinates),
		Coordinates:         coordinates,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
