package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

var (
	hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"method", "path", "status"})

	panicsCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "panics_count",
	}, []string{"method", "path"})

	latency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "request_process_time_seconds",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"method", "path"},
	)
)

func RegisterPrometheusMetrics() {
	prometheus.MustRegister(hits, panicsCount, latency)
}

func RecordHits(method string, path string, statusCode int) {
	hits.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()
}

func RecordPanicsCount(method string, path string) {
	panicsCount.WithLabelValues(method, path).Inc()
}

func RecordLatency(method string, path string, elapsed float64) {
	latency.WithLabelValues(method, path).Observe(elapsed)
}
