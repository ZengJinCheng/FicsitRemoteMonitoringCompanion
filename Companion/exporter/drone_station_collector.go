package exporter

import (
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type DroneStationCollector struct {
	endpoint       string
	metricsDropper *MetricsDropper
}

func NewDroneStationCollector(endpoint string) *DroneStationCollector {
	return &DroneStationCollector{
		endpoint: endpoint,
		metricsDropper: NewMetricsDropper(
			DronePortFuelRate,
			DronePortFuelAmount,
			DronePortRndTrip,
		),
	}
}

func (c *DroneStationCollector) Collect(frmAddress string, sessionName string) {
	details := []DroneStationDetails{}
	err := retrieveData(frmAddress+c.endpoint, &details)
	if err != nil {
		c.metricsDropper.DropStaleMetricLabels()
		log.Printf("从FRM读取无人机站统计数据出错: %s\n", err)
		return
	}

	powerInfo := map[float64]float64{}
	maxPowerInfo := map[float64]float64{}
	for _, d := range details {
		c.metricsDropper.CacheFreshMetricLabel(prometheus.Labels{"url": frmAddress, "session_name": sessionName, "id": d.Id})
		id := d.Id
		home := d.HomeStation
		paired := d.PairedStation

		if len(d.Fuel) > 0 {
			DronePortFuelAmount.WithLabelValues(id, home, d.Fuel[0].Name, frmAddress, sessionName).Set(d.Fuel[0].Amount)
			DronePortFuelRate.WithLabelValues(id, home, d.Fuel[0].Name, frmAddress, sessionName).Set(d.ActiveFuel.Rate)
			DronePortRndTrip.WithLabelValues(id, home, paired, frmAddress, sessionName).Set(d.LatestRndTrip)
		}

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
		DronePortPower.WithLabelValues(cid, frmAddress, sessionName).Set(powerConsumed)
	}

	c.metricsDropper.DropStaleMetricLabels()
}

func (c *DroneStationCollector) DropCache() {}
