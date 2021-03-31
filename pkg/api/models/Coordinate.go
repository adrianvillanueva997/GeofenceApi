package models

type JSONCoordinate struct {
	Description string   `json:"description"`
	Geometry    Geometry `json:"geometry"`
}

type Coordinate struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoJSON struct {
	Type     string   `json:"type"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
