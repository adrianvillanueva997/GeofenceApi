package v1

import (
	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/paulmach/orb/geojson"
)

// package internal variables that are updated whenever there is a polygon update
var polygon models.GeoJSON
var encodedPolygon *geojson.Feature
