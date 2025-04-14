package exporter


type DroneFuelInventory struct {
	Name   string  `json:"Name"`
	Amount float64 `json:"Amount"`
}

type DroneActiveFuel struct {
	Name string  `json:"FuelName"`
	Rate float64 `json:"EstimatedFuelCostRate"`
}

type DroneStationDetails struct {
	Id                     string               `json:"ID"`
	HomeStation            string               `json:"Name"`
	PairedStation          string               `json:"PairedStation"`
	DroneStatus            string               `json:"DroneStatus"`
	AvgIncRate             float64              `json:"AvgIncRate"`
	AvgIncStack            float64              `json:"AvgIncStack"`
	AvgOutRate             float64              `json:"AvgOutRate"`
	AvgOutStack            float64              `json"AvgOutStack"`
	AvgRndTrip             string               `json:"AvgRndTrip"`
	AvgTotalIncRate        float64              `json:"AvgTotalIncRate"`
	AvgTotalIncStack       float64              `json:"AvgTotalIncStack"`
	AvgTotalOutRate        float64              `json:"AvgTotalOutRate"`
	AvgTotalOutStack       float64              `json:"AvgTotalOutStack"`
	AvgTripIncAmt          float64              `json:"AvgTripIncAmt"`
	EstTotalTransRate      float64              `json:"EstTotalTransRate"`
	EstTransRate           float64              `json:"EstTransRate"`
	EstLatestTotalIncStack float64              `json:"EstLatestTotalIncStack"`
	EstLatestTotalOutStack float64              `json:"EstLatestTotalOutStack"`
	LatestIncStack         float64              `json:"LatestIncStack"`
	LatestOutStack         float64              `json:"LatestOutStack"`
	LatestRndTrip          float64              `json:"LatestRndTrip"`
	LatestTripIncAmt       float64              `json:"LatestTripIncAmt"`
	LatestTripOutAmt       float64              `json:"LatestTripOutAmt"`
	MedianRndTrip          string               `json:"MedianRndTrip"`
	MedianTripIncAmt       float64              `json:"MedianTripIncAmt"`
	MedianTripOutAmt       float64              `json:"MedianTripOutAmt"`
	PowerInfo              PowerInfo            `json:"PowerInfo"`
	Fuel                   []DroneFuelInventory `json:"FuelInventory"`
	ActiveFuel             DroneActiveFuel      `json:"ActiveFuel"`
}