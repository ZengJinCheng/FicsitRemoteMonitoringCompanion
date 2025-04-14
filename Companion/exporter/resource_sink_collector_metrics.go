package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ResourceSinkPower = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "resource_sink_power",
		Help: "AWESOME当前消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})

	ResourceSinkPowerMax = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "resource_sink_power_max",
		Help: "AWESOME最大消耗电力，单位为 MW",
	}, []string{
		"circuit_id",
	})

	ResourceSinkTotalPoints = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "resource_sink_total_points",
		Help: "AWESOME 累计积分总数",
	}, []string{
		"sink_type",
	})

	ResourceSinkPointsToCoupon = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "resource_sink_points_to_coupon",
		Help: "AWESOME 下一张优惠劵的积分数",
	}, []string{
		"sink_type",
	})

	ResourceSinkPercent = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "resource_sink_percent",
		Help: "AWESOME 下一张优惠券的进度百分比",
	}, []string{
		"sink_type",
	})

	ResourceSinkCollectedCoupons = RegisterNewGaugeVec(prometheus.GaugeOpts{
		Name: "resource_sink_collected_coupons",
		Help: "AWESOME 里的优惠券数量",
	}, []string{})
)
