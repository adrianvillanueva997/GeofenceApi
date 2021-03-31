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
