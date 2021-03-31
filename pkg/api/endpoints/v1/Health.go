package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

func Health(ctx *fiber.Ctx) error {
	response := models.HealthCheck{
		Status:  fiber.StatusOK,
		Message: "OK",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
