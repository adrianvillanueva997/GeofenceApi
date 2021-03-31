package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
)

func jsonToGeoJSON(jsonModel models.JSONCoordinate) models.GeoJSON {
	geometry := models.Geometry{
		Type:        jsonModel.Geometry.Type,
		Coordinates: jsonModel.Geometry.Coordinates,
	}
	geoJSON := models.GeoJSON{
		Type:     "Feature",
		Geometry: geometry,
	}
	return geoJSON
}

func ReadJSON() models.GeoJSON {
	jsonFile, err := os.Open("data/geofence.json")
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Println("JSON File successfully opened")
	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panicf(err.Error())
	}
	var JSONCoordinates models.JSONCoordinate
	err = json.Unmarshal(byteData, &JSONCoordinates)
	if err != nil {
		log.Panicf(err.Error())
	}
	geoJSON := jsonToGeoJSON(JSONCoordinates)
	return geoJSON
}
