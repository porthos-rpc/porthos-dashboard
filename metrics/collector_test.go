package metrics

import (
	"os"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

func TestCollector(t *testing.T) {
	collector := NewCollector(os.Getenv("BROKER_URL"))
	go collector.Start()

	payload := []byte(`[{"serviceName":"UserService","methodName":"doSomethingThatReturnsValue","responsetime":51996,"statusCode":200},
						{"serviceName":"UserService","methodName":"doSomethingThatReturnsValue","responsetime":51996,"statusCode":200}]`)

	for i := 0; i < 3; i++ {
		collector.channel.Publish(
			"",
			metricsQueueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        payload,
			})
	}

	collectedMetricsCount := 0

	go func() {
		for _ = range collector.MetricsChannel() {
			collectedMetricsCount++
		}
	}()

	<-time.After(2 * time.Second)

	if collectedMetricsCount != 6 {
		t.Errorf("Excepted 6 collected metrics, got %d", collectedMetricsCount)
	}
}
