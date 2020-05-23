package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Struct which contains pointers to all the
// prometheus descriptor defined for this exporter
type AccuweatherCollector struct {
	fooMetric *prometheus.Desc
	barMetric *prometheus.Desc
}

func NewAccuweatherCollector() *AccuweatherCollector {
	return &AccuweatherCollector{
		fooMetric: prometheus.NewDesc("accuweather_foo_metric",
			"Shows weather a foo has occured in the cluster", nil, nil),
		barMetric: prometheus.NewDesc("accuweather_bar_metric",
			"Shows weather a bar has occurred in the cluster", nil, nil),
	}
}

// Each and every collector must implement the Describe function.
// It writes all the descriptors to the prometheus desc channel
func (collector *AccuweatherCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.fooMetric
	ch <- collector.barMetric
}

// Collect implements the required collect function for all
// Prometheus collector
func (collector *AccuweatherCollector) Collect(ch chan<- prometheus.Metric) {
	// Logic to determine proper metric value to return to
	// prometheus. For each descriptor
	var metricValue int64
	metricValue = time.Now().Unix()

	// Write latest value for each metric in the prometheus metric channel
	ch <- prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, float64(metricValue))
	ch <- prometheus.MustNewConstMetric(collector.barMetric, prometheus.GaugeValue, float64(metricValue))
}
