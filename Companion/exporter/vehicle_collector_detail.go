package exporter

type VehicleDetails struct {
	Id           string   `json:"ID"`
	VehicleType  string   `json:"Name"`
	Location     Location `json:"location"`
	ForwardSpeed float64  `json:"ForwardSpeed"`
	AutoPilot    bool     `json:"Autopilot"`
	Fuel         []Fuel   `json:"FuelInventory"`
	PathName     string   `json:"PathName"`
	DepartTime   time.Time
	Departed     bool
	LastTracked  time.Time
}

type Fuel struct {
	Name   string  `json:"Name"`
	Amount float64 `json:"Amount"`
}
