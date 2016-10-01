package models

import "time"

// AggregatedMetric represents metrics that were aggregated during a time span.
type AggregatedMetric struct {
	ServiceName  string        `db:"service_name" json:"service_name"`
	MethodName   string        `db:"methodName" json:"methodName"`
	Timestamp    time.Time     `db:"timestamp" json:"timestamp"`
	Throughput   int           `db:"throughput" json:"throughput"`
	ResponseTime time.Duration `db:"responseTime" json:"responseTime"`
	Status2XX    int           `db:"status2XX" json:"status2XX"`
}
