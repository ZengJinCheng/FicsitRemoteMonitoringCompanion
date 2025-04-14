package exporter

type ProductionDetails struct {
	ItemName           string  `json:"Name"`
	ProdPercent        float64 `json:"ProdPercent"`
	ConsPercent        float64 `json:"ConsPercent"`
	CurrentProduction  float64 `json:"CurrentProd"`
	CurrentConsumption float64 `json:"CurrentConsumed"`
	MaxProd            float64 `json:"MaxProd"`
	MaxConsumed        float64 `json:"MaxConsumed"`
}
