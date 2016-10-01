package metrics

import (
	"fmt"

	"github.com/streadway/amqp"
)

const (
	metricsQueueName = "porthos.metrics"
)

// Collector is responsible for consuming metrics from the broker and aggregate them.
type Collector struct {
	broker          *amqp.Connection
	channel         *amqp.Channel
	deliveryChannel <-chan amqp.Delivery
	metricsChannel  chan *MetricEntry
}

// NewCollector creates a new metrics collector.
func NewCollector(brokerURL string) *Collector {
	broker, err := NewBroker(brokerURL)

	if err != nil {
		panic(err)
	}

	ch, err := broker.Channel()

	if err != nil {
		panic(err)
	}

	_, err = ch.QueueDeclare(
		metricsQueueName, // name
		true,             // durable
		false,            // delete when usused
		false,            // exclusive
		false,            // noWait
		nil,              // arguments
	)

	dc, _ := ch.Consume(
		"porthos.metrics", // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	return &Collector{broker, ch, dc, make(chan *MetricEntry)}
}

// Start the metrics collector.
func (c *Collector) Start() {
	for d := range c.deliveryChannel {
		metrics, err := UnmarshalMetrics(d.Body)

		if err != nil {
			fmt.Errorf("Error parsing metrics %s.", err)
		}

		for _, m := range metrics {
			c.metricsChannel <- m
		}
	}
}

// Stop the metrics collector and release resources.
func (c *Collector) Stop() {
	c.channel.Close()
	c.broker.Close()
}

// MetricsChannel returns a read-only channel of metrics.
func (c *Collector) MetricsChannel() <-chan *MetricEntry {
	return c.metricsChannel
}
