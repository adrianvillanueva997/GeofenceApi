package models

type Coordinate struct {
	Description string `json:"description"`
	Geometry    struct {
		Type        string `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
}
