package models

type Points struct {
	Longitude float64 `json:"longitude" validate:"required,number"`
	Latitude  float64 `json:"latitude" validate:"required,number"`
}

type Update struct {
	ID        int     `json:"id" validate:"required,number"`
	Longitude float64 `json:"longitude" validate:"required,number"`
	Latitude  float64 `json:"latitude" validate:"required,number"`
}
