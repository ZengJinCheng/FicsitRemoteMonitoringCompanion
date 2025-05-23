package exporter

import (
	"log"
	"strconv"
)

type ResourceSinkCollector struct {
	buildingEndpoint          string
	globalResourceEndpoint    string
	globalExplorationEndpoint string
}

func NewResourceSinkCollector(buildingEndpoint, globalResourceEndpoint, globalExplorationEndpoint string) *ResourceSinkCollector {
	return &ResourceSinkCollector{
		buildingEndpoint:          buildingEndpoint,
		globalResourceEndpoint:    globalResourceEndpoint,
		globalExplorationEndpoint: globalExplorationEndpoint,
	}
}

func (c *ResourceSinkCollector) Collect(frmAddress string, sessionName string) {
	buildingDetails := []ResourceSinkDetails{}
	err := retrieveData(frmAddress+c.buildingEndpoint, &buildingDetails)
	if err != nil {
		log.Printf("从FRM读取资源回收器详细统计信息时出错: %s\n", err)
		return
	}

	globalResourceDetails := []GlobalSinkDetails{}
	err = retrieveData(frmAddress+c.globalResourceEndpoint, &globalResourceDetails)
	if err != nil {
		log.Printf("从FRM读取标准点数统计信息时出错: %s\n", err)
		return
	}

	for _, d := range globalResourceDetails {
		ResourceSinkTotalPoints.WithLabelValues(d.SinkType, frmAddress, sessionName).Set(float64(d.TotalPoints))
		ResourceSinkPointsToCoupon.WithLabelValues(d.SinkType, frmAddress, sessionName).Set(float64(d.PointsToCoupon))
		ResourceSinkPercent.WithLabelValues(d.SinkType, frmAddress, sessionName).Set(float64(d.Percent))
		ResourceSinkCollectedCoupons.WithLabelValues(frmAddress, sessionName).Set(float64(d.NumCoupon))
	}

	globalExplorationDetails := []GlobalSinkDetails{}
	err = retrieveData(frmAddress+c.globalExplorationEndpoint, &globalExplorationDetails)
	if err != nil {
		log.Printf("从FRM读取DNA点数统计信息时出错: %s\n", err)
		return
	}

	for _, d := range globalExplorationDetails {
		ResourceSinkTotalPoints.WithLabelValues(d.SinkType, frmAddress, sessionName).Set(float64(d.TotalPoints))
		ResourceSinkPointsToCoupon.WithLabelValues(d.SinkType, frmAddress, sessionName).Set(float64(d.PointsToCoupon))
		ResourceSinkPercent.WithLabelValues(d.SinkType, frmAddress, sessionName).Set(float64(d.Percent))
	}

	powerInfo := map[float64]float64{}
	maxPowerInfo := map[float64]float64{}
	for _, d := range buildingDetails {
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
		ResourceSinkPower.WithLabelValues(cid, frmAddress, sessionName).Set(powerConsumed)
	}
	for circuitId, powerConsumed := range maxPowerInfo {
		cid := strconv.FormatFloat(circuitId, 'f', -1, 64)
		ResourceSinkPowerMax.WithLabelValues(cid, frmAddress, sessionName).Set(powerConsumed)
	}
}

func (c *ResourceSinkCollector) DropCache() {}
