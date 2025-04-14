package exporter

import (
	"time"
)

type TimeTable struct {
	StationName string `json:"StationName"`
}

type TrainCar struct {
	Name           string  `json:"Name"`
	TotalMass      float64 `json:"TotalMass"`
	PayloadMass    float64 `json:"PayloadMass"`
	MaxPayloadMass float64 `json:"MaxPayloadMass"`
}

type TrainDetails struct {
	TrainName        string      `json:"Name"`
	TrainStation     string      `json:"TrainStation"`
	Derailed         bool        `json:"Derailed"`
	Status           string      `json:"Status"` //"Self-Driving",
	TimeTable        []TimeTable `json:"TimeTable"`
	TrainCars        []TrainCar  `json:"Vehicles"`
	PowerInfo        PowerInfo   `json:"PowerInfo"`
	ArrivalTime      time.Time
	StationCounter   int
	FirstArrivalTime time.Time
	LastTracked      time.Time
}