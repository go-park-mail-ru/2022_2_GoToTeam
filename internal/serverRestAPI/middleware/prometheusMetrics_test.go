package middleware

import "testing"

func TestRegisterPrometheusMetrics(t *testing.T) {
	RegisterPrometheusMetrics()
}

func TestRecordHits(t *testing.T) {
	RecordHits("POST", "testStr", 200)
}

func TestRecordPanicsCount(t *testing.T) {
	RecordPanicsCount("POST", "testStr")
}

func TestRecordLatency(t *testing.T) {
	RecordLatency("POST", "testStr", 3.0)
}
