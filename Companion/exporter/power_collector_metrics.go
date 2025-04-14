package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	PowerConsumed = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "power_consumed",
		Help: "选定电路的当前电力消耗",
	}, []string{
		"circuit_id",
	})

	PowerCapacity = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "power_capacity",
		Help: "选定电路的电力理论最大值",
	}, []string{
		"circuit_id",
	})

	PowerProduction = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "power_production",
		Help: "选定电路的当前电力生产量",
	}, []string{
		"circuit_id",
	})

	PowerMaxConsumed = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "power_max_consumed",
		Help: "选定电路的最小消耗电量(需要这么多才能启动)",
	}, []string{
		"circuit_id",
	})

	BatteryDifferential = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "battery_differential",
		Help: "电池组输入或输出的电量盈余或不足。正值表示给电池充电，负值表示消耗电池电量。",
	}, []string{
		"circuit_id",
	})

	BatteryPercent = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "battery_percent",
		Help: "电池组充电百分比",
	}, []string{
		"circuit_id",
	})

	BatteryCapacity = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "battery_capacity",
		Help: "电池组总容量",
	}, []string{
		"circuit_id",
	})

	BatterySecondsEmpty = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "battery_seconds_empty",
		Help: "电池耗尽还有多少秒",
	}, []string{
		"circuit_id",
	})

	BatterySecondsFull = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "battery_seconds_full",
		Help: "电池充满所需秒数",
	}, []string{
		"circuit_id",
	})

	FuseTriggered = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "fuse_triggered",
		Help: "保险丝被触发了吗",
	}, []string{
		"circuit_id",
	})
)
