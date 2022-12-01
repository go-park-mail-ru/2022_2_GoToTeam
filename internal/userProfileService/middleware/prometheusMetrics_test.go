package middleware

import "testing"

func TestRegisterPrometheusMetrics(t *testing.T) {
	RegisterPrometheusMetrics()
}

func TestRecordHits(t *testing.T) {
	RecordHits("testStr")
}

func TestRecordPanicsCount(t *testing.T) {
	RecordPanicsCount("testStr")
}

func TestRecordLatency(t *testing.T) {
	RecordLatency("testStr", 3.0)
}
