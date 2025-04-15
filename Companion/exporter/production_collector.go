package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type ProductionCollector struct {
	endpoint       string
	metricsDropper *MetricsDropper
}

func NewProductionCollector(endpoint string) *ProductionCollector {
	return &ProductionCollector{
		endpoint: endpoint,
		metricsDropper: NewMetricsDropper(
			ItemsProducedPerMin,
			ItemsConsumedPerMin,
			ItemProductionCapacityPercent,
			ItemConsumptionCapacityPercent,
			ItemProductionCapacityPerMinute,
			ItemConsumptionCapacityPerMinute,
		),
	}
}

func (c *ProductionCollector) Collect(frmAddress string, sessionName string) {
	details := []ProductionDetails{}
	err := retrieveData(frmAddress+c.endpoint, &details)
	if err != nil {
		c.metricsDropper.DropStaleMetricLabels()
		log.Printf("从FRM读取生产统计信息时出错: %s\n", err)
		return
	}

	for _, d := range details {
		c.metricsDropper.CacheFreshMetricLabel(prometheus.Labels{"url": frmAddress, "session_name": sessionName, "item_name": d.ItemName})
		ItemsProducedPerMin.WithLabelValues(d.ItemName, frmAddress, sessionName).Set(d.CurrentProduction)
		ItemsConsumedPerMin.WithLabelValues(d.ItemName, frmAddress, sessionName).Set(d.CurrentConsumption)

		ItemProductionCapacityPercent.WithLabelValues(d.ItemName, frmAddress, sessionName).Set(d.ProdPercent)
		ItemConsumptionCapacityPercent.WithLabelValues(d.ItemName, frmAddress, sessionName).Set(d.ConsPercent)
		ItemProductionCapacityPerMinute.WithLabelValues(d.ItemName, frmAddress, sessionName).Set(d.MaxProd)
		ItemConsumptionCapacityPerMinute.WithLabelValues(d.ItemName, frmAddress, sessionName).Set(d.MaxConsumed)
	}
	c.metricsDropper.DropStaleMetricLabels()
}

func (c *ProductionCollector) DropCache() {}
