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

const (
	_DEFAULT_COMPONENT_STRING_FOR_LOGGER  = "unconfigured"
	_DEFAULT_LAYER_STRING_FOR_LOGGER      = "unconfigured"
	_DEFAULT_REQUEST_ID_STRING_FOR_LOGGER = "not request"
	_DEFAULT_USER_EMAIL_STRING_FOR_LOGGER = "unauthorized"
)

type Logger struct {
	mainLogrusLogger *logrus.Logger
	LogrusLogger     *logrus.Entry
	// Another loggers
}

func NewLogger(logLevel string, logFilePath string) (*Logger, error) {
	logrusLogger, err := newLogrusLogger(logLevel, logFilePath)
	if err != nil {
		return nil, fmt.Errorf("can not configure logrus logger: %w", err)
	}

	return &Logger{
		mainLogrusLogger: logrusLogger,
		LogrusLogger: logrusLogger.WithFields(
			logrus.Fields{
				"component": _DEFAULT_COMPONENT_STRING_FOR_LOGGER,
				"layer":     _DEFAULT_LAYER_STRING_FOR_LOGGER,
				"requestId": _DEFAULT_REQUEST_ID_STRING_FOR_LOGGER,
				"userEmail": _DEFAULT_USER_EMAIL_STRING_FOR_LOGGER,
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
		LogFormat:       "[%lvl%][component: %component%][layer: %layer%][requestId: %requestId%][userEmail: %userEmail%][%time%]: %msg%\n_________________________________\n\n",
	}
	logrusLogger.SetFormatter(formatter)

	return logrusLogger, nil
}

func (l *Logger) ConfigureLogger(componentName string, layer string) *Logger {
	return &Logger{
		mainLogrusLogger: l.mainLogrusLogger,
		LogrusLogger: l.LogrusLogger.WithFields(logrus.Fields{
			"component": componentName,
			"layer":     layer,
		}),
	}
}

func (l *Logger) LogrusLoggerWithContext(ctx context.Context) *logrus.Entry {
	requestId := ctx.Value(domain.REQUEST_ID_KEY_FOR_CONTEXT)
	if requestId == nil {
		requestId = _DEFAULT_REQUEST_ID_STRING_FOR_LOGGER
	}
	userEmail := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	if userEmail == nil {
		userEmail = _DEFAULT_USER_EMAIL_STRING_FOR_LOGGER
	}

	return l.LogrusLogger.WithFields(logrus.Fields{
		"requestId": requestId,
		"userEmail": userEmail,
	})
}
