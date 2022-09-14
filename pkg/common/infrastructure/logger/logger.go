package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

type Logger interface {
	WithField(string, interface{}) Logger
	WithFields(Fields) Logger
	Info(...interface{})
	Error(error, ...interface{})
}

type MainLogger interface {
	Logger
	FatalError(error, ...interface{})
}

const appNameKey = "app_name"

type Config struct {
	AppName string
}

func NewLogger(config *Config) MainLogger {
	impl := logrus.New()
	impl.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap:        fieldMap,
	})
	impl.AddHook(&StackTraceHook{})

	return &logger{
		FieldLogger: impl.WithField(appNameKey, config.AppName),
	}
}

type logger struct {
	logrus.FieldLogger
}

var fieldMap = logrus.FieldMap{
	logrus.FieldKeyTime: "@timestamp",
	logrus.FieldKeyMsg:  "message",
}

func (l *logger) WithField(key string, value interface{}) Logger {
	return &logger{l.FieldLogger.WithField(key, value)}
}

func (l *logger) WithFields(fields Fields) Logger {
	return &logger{l.FieldLogger.WithFields(logrus.Fields(fields))}
}

func (l *logger) Error(err error, args ...interface{}) {
	l.FieldLogger.WithError(err).Error(args...)
}

func (l *logger) FatalError(err error, args ...interface{}) {
	l.FieldLogger.WithError(err).Fatal(args...)
}
