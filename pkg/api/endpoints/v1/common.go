package v1

import (
	"time"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
	"github.com/paulmach/orb/geojson"
)

// package internal variables that are updated whenever there is a polygon update or an error occurs
var polygon models.GeoJSON
var encodedPolygon *geojson.Feature
var errorMonitorArray []models.ErrorMonitor

func now() time.Time {
	return time.Now()
}
