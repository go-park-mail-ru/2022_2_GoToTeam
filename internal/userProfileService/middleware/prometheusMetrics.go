package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"method"})

	panicsCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "panics_count",
	}, []string{"method"})

	latency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "request_process_time_seconds",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"method"},
	)
)

func RegisterPrometheusMetrics() {
	prometheus.MustRegister(hits, panicsCount, latency)
}

func RecordHits(method string) {
	hits.WithLabelValues(method).Inc()
}

func RecordPanicsCount(method string) {
	panicsCount.WithLabelValues(method).Inc()
}

func RecordLatency(method string, elapsed float64) {
	latency.WithLabelValues(method).Observe(elapsed)
}
