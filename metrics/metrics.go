package metrics

// Collector is responsible for consuming metrics from the broker and aggregate them.
type Collector struct {
	a int
}

// NewCollector creates a new metrics collector.
func NewCollector(brokerURL string) *Collector {
	return &Collector{}
}
