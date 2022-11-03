package logger

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
)

type Logger struct {
	LogrusLogger *logrus.Logger
	// Another loggers
}

func NewLogger(componentName string, layer string, logLevel, logFilePath string) (*Logger, error) {
	logrusLogger, err := newLogrusLogger(componentName, layer, logLevel, logFilePath)
	if err != nil {
		return nil, fmt.Errorf("cant configure logrus logger: %w", err)
	}

	return &Logger{LogrusLogger: logrusLogger}, nil
}

func newLogrusLogger(componentName string, layer string, logLevel, logFilePath string) (*logrus.Logger, error) {
	logrusLogger := logrus.New()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logrusLogger.SetLevel(level)

	if len(logFilePath) != 0 {
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		logrusLogger.SetOutput(io.MultiWriter(logFile, os.Stdout))

	} else {
		return nil, errors.New("incorrect logFilePath: is empty")
	}

	formatter := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[component: %component%, layer: %layer%][%lvl%][%component%]: %time%: %msg%\n\n____________________\n\n",
	}
	logrusLogger.WithField("component", componentName)
	logrusLogger.WithField("layer", layer)
	logrusLogger.SetFormatter(formatter)

	return logrusLogger, nil
}
