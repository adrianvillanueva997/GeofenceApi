package v1

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	validationv1 "github.com/adrianvillanueva997/GeofenceApi/pkg/api/validation/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func checkLongitude(longitude float64) bool {
	if longitude >= -180 && longitude <= 180 {
		return true
	}
	return false
}

func checkLatitude(latitude float64) bool {
	if latitude >= -90 && latitude <= 90 {
		return true
	}
	return false
}

func checkCoordinates(coordinates models.Update) bool {
	isCoordinate := true
	for i := 0; i < len(coordinates.Polygon.Coordinates[0]) && isCoordinate; i++ {
		if !checkLatitude(coordinates.Polygon.Coordinates[0][i][1]) || !checkLongitude(coordinates.Polygon.Coordinates[0][i][0]) {
			isCoordinate = false
		}
	}
	return isCoordinate
}

func isPolygon(features *geojson.Feature) bool {
	_, isPolygon := features.Geometry.(orb.Polygon)
	return isPolygon
}
func encodeGeoJSON(newPolygon models.GeoJSON) (*geojson.Feature, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(newPolygon)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	features, err := geojson.UnmarshalFeature(buffer.Bytes())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	encodedPolygon = features
	return features, err
}

func UpdatePolygon(ctx *fiber.Ctx) error {
	var body models.Update
	err := ctx.BodyParser(&body)
	if err != nil {
		errMessage := models.ErrorResponse{
			Status:       fiber.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	errors := validationv1.ValidateUpdateRequest(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	errors = validationv1.ValidateGeometryStruct(body.Polygon)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if checkCoordinates(body) {
		geojsonData := models.GeoJSON{
			Description: body.Description,
			Type:        "Feature",
			Geometry:    body.Polygon,
		}
		features, err := encodeGeoJSON(geojsonData)
		if err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		if isPolygon(features) {
			polygon.Type = "Feature"
			polygon.Geometry.Type = body.Polygon.Type
			polygon.Description = body.Description
			polygon.Geometry.Coordinates = body.Polygon.Coordinates
			response := models.UpdateCoordinateResponse{
				Status:  fiber.StatusOK,
				Message: "Polygon updated successfully",
			}
			return ctx.Status(response.Status).JSON(response)
		}
		errMessage := models.ErrorResponse{
			Status:       fiber.StatusBadRequest,
			ErrorMessage: "Bad polygon",
		}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON("fail")
}
