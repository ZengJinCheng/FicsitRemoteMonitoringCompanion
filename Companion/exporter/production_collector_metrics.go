package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	ItemsProducedPerMin = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "items_produced_per_min",
		Help: "物品当前每分钟产量",
	}, []string{
		"item_name",
	})
	
	ItemsConsumedPerMin = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "items_consumed_per_min",
		Help: "物品当前消耗量",
	}, []string{
		"item_name",
	})
	
	ItemProductionCapacityPercent = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "item_production_capacity_pc",
		Help: "物品每分钟生产效率百分比",
	}, []string{
		"item_name",
	})

	ItemConsumptionCapacityPercent = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "item_consumption_capacity_pc",
		Help: "物品每分钟消耗效率百分比",
	}, []string{
		"item_name",
	})

	ItemProductionCapacityPerMinute = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "item_production_capacity_per_min",
		Help: "物品理论最大产量",
	}, []string{
		"item_name",
	})
	
	ItemConsumptionCapacityPerMinute = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "item_consumption_capacity_per_min",
		Help: "物品理论最大消耗量",
	}, []string{
		"item_name",
	})

)
