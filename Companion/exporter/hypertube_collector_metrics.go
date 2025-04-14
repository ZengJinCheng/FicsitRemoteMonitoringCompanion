package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HypertubePower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "hypertube_power",
		Help: "hypertube power use in MW",
	}, []string{
		"circuit_id",
	})

	HypertubePowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "hypertube_power_max",
		Help: "hypertube max power use in MW",
	}, []string{
		"circuit_id",
	})
)