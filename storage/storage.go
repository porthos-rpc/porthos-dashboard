package storage

import (
	"time"

	"github.com/porthos-rpc/porthos-dashboard/models"
)

// Storage structure.
type Storage interface {
	InsertAggregatedMetric(aggregatedMetric *models.AggregatedMetric)
	FindMethodMetrics(since time.Time) ([]*models.ServiceMethodMetrics, error)
}
