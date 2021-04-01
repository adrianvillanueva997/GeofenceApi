package v1

import (
	"fmt"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	validationv1 "github.com/adrianvillanueva997/GeofenceApi/pkg/api/validation/v1"
	"github.com/gofiber/fiber/v2"
)

// Checks if the ID provided is less than the total amount of ids in the coordinates array
func checkPoints(id int) bool {
	if len(coordinates.Geometry.Coordinates[0]) >= id {
		return true
	}
	return false
}
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

func UpdatePoint(ctx *fiber.Ctx) error {
	var body models.Update
	err := ctx.BodyParser(&body)
	if err != nil {
		errMessage := models.ErrorResponse{
			Status:       fiber.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	fmt.Println(body)
	errors := validationv1.ValidateUpdateRequest(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if checkLatitude(body.Latitude) == false {
		errMessage := models.ErrorResponse{Status: fiber.StatusBadRequest, ErrorMessage: "Bad Latitude"}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	if checkLongitude(body.Longitude) == false {
		errMessage := models.ErrorResponse{Status: fiber.StatusBadRequest, ErrorMessage: "Bad Longitude"}
		return ctx.Status(errMessage.Status).JSON(errMessage)
	}
	if checkPoints(body.ID) {
		coordinates.Geometry.Coordinates[0][body.ID][1] = body.Longitude
		coordinates.Geometry.Coordinates[0][body.ID][0] = body.Latitude
		response := models.UpdateCoordinateResponse{
			Status:  fiber.StatusOK,
			Message: "Coordinates updated succesfully",
		}
		return ctx.Status(response.Status).JSON(response)
	}
	errMessage := models.ErrorResponse{
		Status:       fiber.StatusBadRequest,
		ErrorMessage: "Bad Coordinates ID",
	}
	return ctx.Status(errMessage.Status).JSON(errMessage)
}
