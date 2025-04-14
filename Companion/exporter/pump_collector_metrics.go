package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PumpPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "pump_power",
		Help: "水泵当前消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})

	PumpPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "pump_power_max",
		Help: "水泵最大消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})
)