package logger

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	logger, err := NewLogger("DEBUG", "../../../../configs/authSessionService/server.toml")
	if err != nil {
		t.Error(err)
	}
	logger = logger.ConfigureLogger("componentA", "delivery")
	entry := logger.LogrusLoggerWithContext(context.Background())
	if entry == nil {
		t.Error(err)
	}
}

func TestLoggerNegative(t *testing.T) {
	_, err := NewLogger("asdasdasd", "../../../../configs/authSessionService/server.toml")
	if err == nil {
		t.Error(err)
	}
	_, err = NewLogger("DEBUG", "")
	if err == nil {
		t.Error(err)
	}
}
