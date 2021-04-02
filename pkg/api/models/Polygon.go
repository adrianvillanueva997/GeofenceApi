package models

type JSONCoordinate struct {
	Description string   `json:"description"`
	Geometry    Geometry `json:"geometry"`
}

type GeoJSON struct {
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Geometry    Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
