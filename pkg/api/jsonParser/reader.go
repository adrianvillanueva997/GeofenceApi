package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
)

func ReadJSON() []models.Coordinate {
	jsonFile, err := os.Open("data/geofence.json")
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Println("JSON File successfully opened")
	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panicf(err.Error())
	}
	var JSONcoordinates models.JSONCoordinate
	err = json.Unmarshal(byteData, &JSONcoordinates)
	if err != nil {
		log.Panicf(err.Error())
	}
	coordinates := jsonToCoordinates(JSONcoordinates)
	return coordinates
}

func jsonToCoordinates(jSONCoordinates models.JSONCoordinate) []models.Coordinate {
	var coordinates []models.Coordinate
	for i := 0; i < len(jSONCoordinates.Geometry.Coordinates[0]); i++ {
		var coordinate models.Coordinate
		coordinate.ID = i
		coordinate.Latitude = jSONCoordinates.Geometry.Coordinates[0][i][0]
		coordinate.Longitude = jSONCoordinates.Geometry.Coordinates[0][i][1]
		coordinates = append(coordinates, coordinate)
	}
	return coordinates
}
