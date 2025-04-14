package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DronePortFuelRate = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "drone_port_fuel_rate",
		Help: "无人机平台燃料消耗率",
	}, []string{
		"id",
		"home_station",
		"fuel_name",
	})
	DronePortFuelAmount = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "drone_port_fuel_amount",
		Help: "无人机平台库存中的燃料量",
	}, []string{
		"id",
		"home_station",
		"fuel_name",
	})
	DronePortRndTrip = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "drone_port_round_trip_seconds",
		Help: "记录无人机往返时间（秒）",
	}, []string{
		"id",
		"home_station",
		"paired_station",
	})
	DronePortPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "drone_port_power",
		Help: "无人机平台当前消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})
	DronePortPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "drone_port_power_max",
		Help: "无人机平台最大消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})
)
