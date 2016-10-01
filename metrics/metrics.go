package metrics

import (
	"encoding/json"
	"fmt"
	"time"
)

// MetricEntry represents a metric of a RPC method call.
type MetricEntry struct {
	ServiceName  string        `json:"serviceName"`
	MethodName   string        `json:"methodName"`
	ResponseTime time.Duration `json:"responsetime"`
	StatusCode   int16         `json:"statusCode"`
}

// Key returns this metric key as "ServiceName.methodName".
func (m *MetricEntry) Key() string {
	return fmt.Sprintf("%s.%s", m.ServiceName, m.MethodName)
}

// UnmarshalMetrics takes an array of bytes (json encoded), allocates and returns metrics array.
func UnmarshalMetrics(bytes []byte) ([]*MetricEntry, error) {
	var metrics []*MetricEntry
	err := json.Unmarshal(bytes, &metrics)

	return metrics, err
}
