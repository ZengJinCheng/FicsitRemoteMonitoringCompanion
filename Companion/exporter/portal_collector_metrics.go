package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PortalPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "portal_power",
		Help: "传送门当前消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})

	PortalPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "portal_power_max",
		Help: "传送门最大消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})
)