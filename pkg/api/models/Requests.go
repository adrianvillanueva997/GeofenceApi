package models

type Points struct {
	Longitude float64 `json:"longitude" validate:"required,number"`
	Latitude  float64 `json:"latitude" validate:"required,number"`
}

type Update struct {
	Description string   `json:"description" validate:"required"`
	Polygon     Geometry `json:"geometry" validate:"required"`
}
