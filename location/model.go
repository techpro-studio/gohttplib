package location

import (
	"encoding/json"
)

type LocationParameters struct {
	Latitude    float64
	Longitude   float64
	MinDistance int64
	MaxDistance int64
}

func NewObjectLocation(latitude, longitude float64) *ObjectLocation {
	return &ObjectLocation{"Point", []float64{longitude, latitude}}
}

type ObjectLocation struct {
	Type        string
	Coordinates []float64
}

func (self ObjectLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		Latitude:  self.Coordinates[1],
		Longitude: self.Coordinates[0],
	})
}
