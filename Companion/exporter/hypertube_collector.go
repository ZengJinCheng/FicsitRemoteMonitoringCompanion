package exporter

import (
	"log"
	"strconv"
)

type HypertubeCollector struct {
	endpoint string
}

func NewHypertubeCollector(endpoint string) *HypertubeCollector {
	return &HypertubeCollector{
		endpoint: endpoint,
	}
}

func (c *HypertubeCollector) Collect(frmAddress string, sessionName string) {
	details := []HypertubeDetails{}
	err := retrieveData(frmAddress+c.endpoint, &details)
	if err != nil {
		log.Printf("从FRM读取超级管道统计数据出错: %s\n", err)
		return
	}

	powerInfo := map[float64]float64{}
	maxPowerInfo := map[float64]float64{}
	for _, d := range details {
		val, ok := powerInfo[d.PowerInfo.CircuitGroupId]
		if ok {
			powerInfo[d.PowerInfo.CircuitGroupId] = val + d.PowerInfo.PowerConsumed
		} else {
			powerInfo[d.PowerInfo.CircuitGroupId] = d.PowerInfo.PowerConsumed
		}
		val, ok = maxPowerInfo[d.PowerInfo.CircuitGroupId]
		if ok {
			maxPowerInfo[d.PowerInfo.CircuitGroupId] = val + d.PowerInfo.MaxPowerConsumed
		} else {
			maxPowerInfo[d.PowerInfo.CircuitGroupId] = d.PowerInfo.MaxPowerConsumed
		}
	}
	for circuitId, powerConsumed := range powerInfo {
		cid := strconv.FormatFloat(circuitId, 'f', -1, 64)
		HypertubePower.WithLabelValues(cid, frmAddress, sessionName).Set(powerConsumed)
	}
	for circuitId, powerConsumed := range maxPowerInfo {
		cid := strconv.FormatFloat(circuitId, 'f', -1, 64)
		HypertubePowerMax.WithLabelValues(cid, frmAddress, sessionName).Set(powerConsumed)
	}
}

func (c *HypertubeCollector) DropCache() {}
