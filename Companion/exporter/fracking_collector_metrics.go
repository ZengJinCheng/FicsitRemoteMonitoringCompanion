package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	FrackingPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "fracking_power",
		Help: "fracking power use in MW",
	}, []string{
		"circuit_id",
	})

	FrackingPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "fracking_power_max",
		Help: "fracking max power use in MW",
	}, []string{
		"circuit_id",
	})
)