package logger

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
	"path"
)

var globalLogrusLogger *logrus.Logger = nil

type Logger struct {
	globalLogrusLogger *logrus.Logger
	LogrusLogger       *logrus.Entry
	// Another loggers
}

func NewLogger(componentName string, layer string, logLevel, logFilePath string) (*Logger, error) {
	if globalLogrusLogger == nil {
		var err error
		globalLogrusLogger, err = newLogrusLogger(logLevel, logFilePath)
		if err != nil {
			return nil, fmt.Errorf("can not configure logrus logger: %w", err)
		}
	}

	return &Logger{
		globalLogrusLogger: globalLogrusLogger,
		LogrusLogger: globalLogrusLogger.WithFields(
			logrus.Fields{
				"component": componentName,
				"layer":     layer,
			},
		),
	}, nil
}

func newLogrusLogger(logLevel, logFilePath string) (*logrus.Logger, error) {
	logrusLogger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logrusLogger.SetLevel(level)

	if len(logFilePath) == 0 {
		return nil, errors.New("incorrect logFilePath: is empty")
	}

	if err := os.MkdirAll(path.Dir(logFilePath), 0750); err != nil && !os.IsExist(err) {
		return nil, err
	}
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logrusLogger.SetOutput(io.MultiWriter(logFile, os.Stdout)) // Logging to console and logFile

	formatter := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%][component: %component%][layer: %layer%][requestId: %requestId%][%time%]: %msg%\n_________________________________\n\n",
	}
	logrusLogger.SetFormatter(formatter)

	return logrusLogger, nil
}

func (l *Logger) LogrusLoggerWithContext(ctx context.Context) *logrus.Entry {
	return l.LogrusLogger.WithFields(logrus.Fields{
		"requestId": ctx.Value(domain.REQUEST_ID_KEY_FOR_CONTEXT),
	})
}
