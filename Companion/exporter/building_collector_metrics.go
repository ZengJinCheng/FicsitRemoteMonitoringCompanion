package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MachineItemsProducedPerMin = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "machine_items_produced_per_min",
		Help: "建筑物当前每分钟生产多少物品",
	}, []string{
		"item_name",
		"machine_name",
		"x",
		"y",
		"z",
	})

	MachineItemsProducedEffiency = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "machine_items_produced_pc",
		Help: "建筑物生产物品的效率",
	}, []string{
		"item_name",
		"machine_name",
		"x",
		"y",
		"z",
	})
	FactoryPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "factory_power",
		Help: "工厂机器的电力消耗，单位为 MW。不包括提取设备。",
	}, []string{
		"circuit_id",
	})

	FactoryPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "factory_power_max",
		Help: "工厂机器的最大电力消耗，单位为 MW。不包括提取器。",
	}, []string{
		"circuit_id",
	})
)
