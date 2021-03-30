package jsonParser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/adrianvillanueva997/GeofenceApi/pkg/api/models"
)

func ReadJSON() models.Coordinate {
	jsonFile, err := os.Open("data/geofence.json")
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Println("JSON File successfully opened")
	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panicf(err.Error())
	}
	var coordindates models.Coordinate
	err = json.Unmarshal(byteData, &coordindates)
	if err != nil {
		log.Panicf(err.Error())
	}
	return coordindates
}
