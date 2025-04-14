package exporter


type ResourceSinkDetails struct {
	Location  Location  `json:"location"`
	PowerInfo PowerInfo `json:"PowerInfo"`
}

type GlobalSinkDetails struct {
	SinkType       string  `json:"Name"`
	NumCoupon      int     `json:"NumCoupon"`
	TotalPoints    int     `json:"TotalPoints"`
	PointsToCoupon int     `json:"PointsToCoupon"`
	Percent        float64 `json:"Percent"`
}