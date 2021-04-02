package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	validationv1 "github.com/adrianvillanueva997/GeofenceApi/pkg/api/validation/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func isInside(features *geojson.Feature, point orb.Point) bool {
	newPolygon, _ := features.Geometry.(orb.Polygon)
	return planar.PolygonContains(newPolygon, point)
}

func DeviceStatus(ctx *fiber.Ctx) error {
	var body models.Points
	if err := ctx.BodyParser(&body); err != nil {
		errorMessage := models.ErrorResponse{
			Status:       fiber.StatusBadRequest,
			ErrorMessage: "Missing fields in the request",
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(errorMessage)
	}
	errors := validationv1.ValidatePointRequest(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if !checkLatitude(body.Latitude) {
		errMessage := models.ErrorResponse{
			Status:       fiber.StatusBadRequest,
			ErrorMessage: "Latitude not valid",
		}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	if !checkLongitude(body.Longitude) {
		errMessage := models.ErrorResponse{
			Status:       fiber.StatusBadRequest,
			ErrorMessage: "Longitude not valid",
		}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	deviceStatus := models.DeviceStatusResponse{
		Status:  fiber.StatusOK,
		Message: false,
	}
	if encodedPolygon == nil {
		polygonNotLoadedMessage := models.ErrorResponse{
			Status:       fiber.StatusInternalServerError,
			ErrorMessage: "Polygon not loaded yet.",
		}
		return ctx.Status(polygonNotLoadedMessage.Status).JSON(polygonNotLoadedMessage)
	}
	if isInside(encodedPolygon, orb.Point{body.Longitude, body.Latitude}) {
		deviceStatus.Message = true
		return ctx.Status(deviceStatus.Status).JSON(deviceStatus)
	}
	return ctx.Status(deviceStatus.Status).JSON(deviceStatus)
}
