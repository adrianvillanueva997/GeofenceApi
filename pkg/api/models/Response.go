package models

type HealthCheck struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type PolygonResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Polygon GeoJSON `json:"geometry"`
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
