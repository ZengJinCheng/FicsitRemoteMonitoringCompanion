package exporter


type VehicleStationDetails struct {
	Name      string    `json:"Name"`
	Location  Location  `json:"location"`
	PowerInfo PowerInfo `json:"PowerInfo"`
}