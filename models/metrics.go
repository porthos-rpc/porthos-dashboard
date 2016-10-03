package models

import "time"

// AggregatedMetric represents metrics that were aggregated during a time span.
type AggregatedMetric struct {
	ServiceName  string        `db:"serviceName" json:"serviceName"`
	MethodName   string        `db:"methodName" json:"methodName"`
	Timestamp    time.Time     `db:"timestamp" json:"timestamp"`
	Throughput   int           `db:"throughput" json:"throughput"`
	ResponseTime time.Duration `db:"responseTime" json:"responseTime"`
	Status2XX    int           `db:"status2XX" json:"status2XX"`
}

// ServiceMethodMetrics groups all aggregated metric by service.method.
type ServiceMethodMetrics struct {
	ServiceName string `db:"serviceName" json:"serviceName"`
	MethodName  string `db:"methodName" json:"methodName"`

	MinThroughput int `db:"minThroughput" json:"minThroughput"`
	MaxThroughput int `db:"maxThroughput" json:"maxThroughput"`
	AvgThroughput int `db:"avgThroughput" json:"avgThroughput"`

	MinResponseTime time.Duration `db:"minResponseTime" json:"minResponseTime"`
	MaxResponseTime time.Duration `db:"maxResponseTime" json:"maxResponseTime"`
	AvgResponseTime time.Duration `db:"avgResponseTime" json:"avgResponseTime"`

	MinStatus2XX time.Duration `db:"minStatus2XX" json:"minStatus2XX"`
	MaxStatus2XX time.Duration `db:"maxStatus2XX" json:"maxStatus2XX"`
	AvgStatus2XX time.Duration `db:"avgStatus2XX" json:"avgStatus2XX"`
}
