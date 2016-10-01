package metrics

import (
	"sync"
	"time"
)

const (
	metricThroughput   = iota
	metricResponseTime = iota
	metric2XXStatus    = iota
)

// AggregatedMetric represents metrics that were aggregated during a time span.
type AggregatedMetric struct {
	ServiceName  string
	MethodName   string
	Throughput   int
	ResponseTime time.Duration
	Status2XX    int
}

// Aggregator of metrics.
type Aggregator struct {
	metricsChannel           <-chan *MetricEntry
	methods                  map[string]*AggregatedMetric
	mutex                    *sync.Mutex
	ticker                   *time.Ticker
	aggregatedMetricsChannel chan *AggregatedMetric
}

// NewAggregator allocates a new metrics aggregator.
func NewAggregator(metricsChannel <-chan *MetricEntry, shipperInterval time.Duration) *Aggregator {
	methods := make(map[string]*AggregatedMetric)
	mutex := new(sync.Mutex)
	ticker := time.NewTicker(shipperInterval)
	channel := make(chan *AggregatedMetric)

	return &Aggregator{metricsChannel, methods, mutex, ticker, channel}
}

// Start the aggregator.
func (a *Aggregator) Start() {
	for m := range a.metricsChannel {
		a.aggregateMetric(m)
	}
}

// StartShipper the aggregated metrics.
func (a *Aggregator) StartShipper() {
	for range a.ticker.C {
		a.ship()
		a.reset()
	}
}

// AggregatedMetricsChannel returns a read-only metrics channel.
func (a *Aggregator) AggregatedMetricsChannel() <-chan *AggregatedMetric {
	return a.aggregatedMetricsChannel
}

func (a *Aggregator) aggregateMetric(m *MetricEntry) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	key := m.Key()

	if _, ok := a.methods[key]; !ok {
		a.methods[key] = &AggregatedMetric{
			ServiceName:  m.ServiceName,
			MethodName:   m.MethodName,
			Throughput:   0,
			ResponseTime: 0,
			Status2XX:    0,
		}
	}

	a.methods[key].Throughput++
	a.methods[key].ResponseTime += m.ResponseTime
	a.methods[key].Status2XX += getStatusCode2XXFactor(m.StatusCode)
}

func (a *Aggregator) ship() {
	for _, aggregated := range a.methods {
		a.aggregatedMetricsChannel <- aggregated
	}
}

func (a *Aggregator) reset() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.methods = make(map[string]*AggregatedMetric)
}

func getStatusCode2XXFactor(statusCode int16) int {
	if statusCode >= 200 && statusCode <= 299 {
		return 1
	}

	return 0
}
