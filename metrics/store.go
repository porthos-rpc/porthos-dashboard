package metrics

import (
	"github.com/porthos-rpc/porthos-dashboard/models"
	"github.com/porthos-rpc/porthos-dashboard/storage"
)

// StoreAggregatedMetrics starts a aggregated metrics store
func StoreAggregatedMetrics(s storage.Storage, aggregatedMetricsChan <-chan *models.AggregatedMetric) {
	for am := range aggregatedMetricsChan {
		s.InsertAggregatedMetric(am)
	}
}
