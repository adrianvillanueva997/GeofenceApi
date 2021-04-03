package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

func Monitoring(ctx *fiber.Ctx) error {
	monitoringResponse := models.ErrorMonitorResponse{
		Status:        fiber.StatusOK,
		ErrorMessages: errorMonitorArray,
	}
	return ctx.Status(monitoringResponse.Status).JSON(monitoringResponse)
}
