package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPoints(ctx *fiber.Ctx) error {
	coordinatesModel := jsonToCoordinates(coordinates)
	response := models.GetCoordinates{
		Status:              fiber.StatusOK,
		NumberOfCoordinates: len(coordinatesModel),
		Coordinates:         coordinatesModel,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func jsonToCoordinates(jSONCoordinates models.GeoJSON) []models.Coordinate {
	var coords []models.Coordinate
	for i := 0; i < len(jSONCoordinates.Geometry.Coordinates[0]); i++ {
		var coordinate models.Coordinate
		coordinate.ID = i
		coordinate.Latitude = jSONCoordinates.Geometry.Coordinates[0][i][0]
		coordinate.Longitude = jSONCoordinates.Geometry.Coordinates[0][i][1]
		coords = append(coords, coordinate)
	}
	return coords
}
