package exporter


type CargoPlatform struct {
	LoadingDock   string    `json:"LoadingDock"`
	TransferRate  float64   `json:"TransferRate"`
	LoadingStatus string    `json:"LoadingStatus"` // Idle, Loading, Unloading
	LoadingMode   string    `json:"LoadingMode"`
	PowerInfo     PowerInfo `json:"PowerInfo"`
}

type TrainStationDetails struct {
	Name           string          `json:"Name"`
	Location       Location        `json:"location"`
	CargoPlatforms []CargoPlatform `json:"CargoInventory"`
	PowerInfo      PowerInfo       `json:"PowerInfo"`
}