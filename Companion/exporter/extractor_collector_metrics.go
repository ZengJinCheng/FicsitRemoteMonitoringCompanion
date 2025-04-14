package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ExtractorPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "extractor_power",
		Help: "提取器当前消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})

	ExtractorPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "extractor_power_max",
		Help: "提取器最大消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})
)