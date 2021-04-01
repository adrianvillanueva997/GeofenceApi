package v1

import (
	"bytes"
	"encoding/json"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	validationv1 "github.com/adrianvillanueva997/GeofenceApi/pkg/api/validation/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func isInside(features *geojson.Feature, point orb.Point) bool {
	polygon, isPolygon := features.Geometry.(orb.Polygon)
	if isPolygon {
		return planar.PolygonContains(polygon, point)
	}
	return false
}

func encodeGeoJSON() (*geojson.Feature, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(coordinates)
	if err != nil {
		return nil, err
	}
	features, err := geojson.UnmarshalFeature(buffer.Bytes())
	if err != nil {
		return nil, err
	}
	return features, err
}

func CheckDeviceInsidePoint(ctx *fiber.Ctx) error {
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
	features, err := encodeGeoJSON()
	if err != nil {
		errorMessage := models.ErrorResponse{
			Status:       fiber.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorMessage)
	}
	objectStatus := isInside(features, orb.Point{body.Longitude, body.Latitude})
	statusResponse := models.DeviceStatusResponse{
		Status:  fiber.StatusOK,
		Message: objectStatus,
	}
	return ctx.Status(fiber.StatusOK).JSON(statusResponse)
}
