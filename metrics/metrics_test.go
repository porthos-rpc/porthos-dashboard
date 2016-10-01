package metrics

import "testing"

func TestUnmarshalMetrics(t *testing.T) {
	metrics, err := UnmarshalMetrics([]byte("[{\"serviceName\":\"UserService\",\"methodName\":\"doSomethingThatReturnsValue\",\"responsetime\":51996,\"statusCode\":200}, " +
		"{\"serviceName\":\"UserService\",\"methodName\":\"doSomethingThatReturnsValue\",\"responsetime\":51996,\"statusCode\":200}]"))

	if err != nil {
		t.Errorf("Error unmarshlling bytes to json %s", err)
	}

	if len(metrics) != 2 {
		t.Errorf("Excepted 2 metric, got %d", len(metrics))
	}
}

func TestUnmarshalMetricsError(t *testing.T) {
	_, err := UnmarshalMetrics([]byte("[\"serviceNames\":\"UserService\",\"methodName\":\"doSomethingThatReturnsValue\",\"responsetime\":51996,\"statusCode\":200}, " +
		"{\"serviceName\":\"UserService\",\"methodNam\":\"doSomethingThatReturnsValue\",\"responsetime\":51996,\"statusCode\":200}]"))

	if err == nil {
		t.Errorf("Expected error in unmarshal operation")
	}
}
