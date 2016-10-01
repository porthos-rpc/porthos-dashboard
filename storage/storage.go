package storage

import "github.com/porthos-rpc/porthos-dashboard/models"

// Storage structure.
type Storage interface {
	InsertAggregatedMetric(aggregatedMetric *models.AggregatedMetric)
}
