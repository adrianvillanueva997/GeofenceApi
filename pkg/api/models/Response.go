package models

type HealthCheck struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type GetCoordinates struct {
	Status              int          `json:"status"`
	NumberOfCoordinates int          `json:"number_of_coordinates"`
	Coordinates         []Coordinate `json:"coordinates"`
}
type DeviceStatusResponse struct {
	Status  int  `json:"status"`
	Message bool `json:"message"`
}

type UpdateCoordinateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type FieldErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type ErrorResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}
