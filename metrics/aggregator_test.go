package metrics

import (
	"testing"
	"time"
)

func TestAggregateMetric(t *testing.T) {
	metricsChannel := make(chan *MetricEntry)

	a := NewAggregator(metricsChannel, time.Second*1)

	a.aggregateMetric(&MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod",
		ResponseTime: 10,
		StatusCode:   200,
	})

	a.aggregateMetric(&MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod",
		ResponseTime: 20,
		StatusCode:   500,
	})

	a.aggregateMetric(&MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod2",
		ResponseTime: 200,
		StatusCode:   201,
	})

	if len(a.methods) != 2 {
		t.Errorf("Expected 2 aggregated methods, got %d", len(a.methods))
	}

	if a.methods["SampleService.sampleMethod"].ServiceName != "SampleService" {
		t.Errorf("Expected 'SampleService' as service name, got %s", a.methods["SampleService.sampleMethod"].ServiceName)
	}

	if a.methods["SampleService.sampleMethod"].MethodName != "sampleMethod" {
		t.Errorf("Expected 'sampleMethod' as method name, got %s", a.methods["sampleMethod.sampleMethod"].MethodName)
	}

	if a.methods["SampleService.sampleMethod"].Throughput != 2 {
		t.Errorf("Expected 2 as throughput, got %d", a.methods["SampleService.sampleMethod"].Throughput)
	}

	if a.methods["SampleService.sampleMethod"].ResponseTime != 30 {
		t.Errorf("Expected 30 as response time, got %d", a.methods["SampleService.sampleMethod"].ResponseTime)
	}

	if a.methods["SampleService.sampleMethod"].Status2XX != 1 {
		t.Errorf("Expected 1 as status 2XX, got %d", a.methods["SampleService.sampleMethod"].Status2XX)
	}

	if a.methods["SampleService.sampleMethod2"].ServiceName != "SampleService" {
		t.Errorf("Expected 'SampleService' as service name, got %s", a.methods["SampleService.sampleMethod2"].ServiceName)
	}

	if a.methods["SampleService.sampleMethod2"].MethodName != "sampleMethod2" {
		t.Errorf("Expected 'sampleMethod' as method name, got %s", a.methods["sampleMethod.sampleMethod2"].MethodName)
	}

	if a.methods["SampleService.sampleMethod2"].Throughput != 1 {
		t.Errorf("Expected 1 as throughput, got %d", a.methods["SampleService.sampleMethod2"].Throughput)
	}

	if a.methods["SampleService.sampleMethod2"].ResponseTime != 200 {
		t.Errorf("Expected 200 as response time, got %d", a.methods["SampleService.sampleMethod2"].ResponseTime)
	}

	if a.methods["SampleService.sampleMethod2"].Status2XX != 1 {
		t.Errorf("Expected 1 as status 2XX, got %d", a.methods["SampleService.sampleMethod2"].Status2XX)
	}
}

func TestShipper(t *testing.T) {
	metricsChannel := make(chan *MetricEntry)

	a := NewAggregator(metricsChannel, time.Second*2)
	go a.Start()
	go a.StartShipper()

	shippedMetrics := 0

	go func() {
		for range a.AggregatedMetricsChannel() {
			shippedMetrics++
		}
	}()

	metricsChannel <- &MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod",
		ResponseTime: 10,
		StatusCode:   200,
	}

	metricsChannel <- &MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod",
		ResponseTime: 20,
		StatusCode:   500,
	}

	metricsChannel <- &MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod2",
		ResponseTime: 200,
		StatusCode:   201,
	}

	<-time.After(3 * time.Second)

	if shippedMetrics != 2 {
		t.Errorf("Excepted 2 shipped metrics, got %d", shippedMetrics)
	}

	metricsChannel <- &MetricEntry{
		ServiceName:  "SampleService",
		MethodName:   "sampleMethod2",
		ResponseTime: 200,
		StatusCode:   201,
	}

	<-time.After(3 * time.Second)

	if shippedMetrics != 3 {
		t.Errorf("Excepted 3 shipped metrics, got %d", shippedMetrics)
	}
}
