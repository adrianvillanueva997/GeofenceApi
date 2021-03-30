package models

type HealthCheck struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
