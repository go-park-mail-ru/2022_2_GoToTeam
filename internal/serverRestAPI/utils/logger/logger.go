package logger

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
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

	if len(logFilePath) != 0 {
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}

		logrusLogger.SetOutput(io.MultiWriter(logFile, os.Stdout)) // Logging to console and logFile
	} else {
		return nil, errors.New("incorrect logFilePath: is empty")
	}

	formatter := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%][component: %component%][layer: %layer%][requestId: %requestId%][%time%]: %msg%\n____________________\n\n",
	}
	logrusLogger.SetFormatter(formatter)

	return logrusLogger, nil
}
