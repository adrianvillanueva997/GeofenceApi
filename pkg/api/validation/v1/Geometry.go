package validationv1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/go-playground/validator/v10"
)

func ValidateGeometryStruct(structToValidate models.Geometry) []*models.FieldErrorResponse {
	var errors []*models.FieldErrorResponse
	validate := validator.New()
	err := validate.Struct(structToValidate)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.FieldErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
