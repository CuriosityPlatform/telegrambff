package logger

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type StackTraceHook struct{}

const keyStack = "stack"

func (h *StackTraceHook) Fire(entry *logrus.Entry) error {
	val, ok := entry.Data[logrus.ErrorKey]
	if !ok {
		return nil
	}

	err, ok := val.(error)
	if !ok {
		return nil
	}

	if err == nil {
		delete(entry.Data, logrus.ErrorKey)
		return nil
	}

	trace := GetStackTrace(err)
	if trace != nil {
		entry.Data[keyStack] = trace
	}

	entry.Data[logrus.ErrorKey] = err.Error()

	return nil
}

func (h *StackTraceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func GetStackTrace(err error) []string {
	tracer := getOldestStackTracer(err)
	if tracer == nil {
		return nil
	}

	stackTrace := tracer.StackTrace()
	return toStringSlice(stackTrace)
}

func toStringSlice(stackTrace errors.StackTrace) (result []string) {
	for _, f := range stackTrace {
		result = append(result, fmt.Sprintf("%+v", f))
	}
	return
}

func getOldestStackTracer(err error) (oldestStackTracer stackTracer) {
	for err != nil {
		if tracer, ok := err.(stackTracer); ok {
			oldestStackTracer = tracer
		}

		err = getCause(err)
	}

	return
}

type causer interface {
	Cause() error
}

func getCause(err error) error {
	if c, ok := err.(causer); ok {
		return c.Cause()
	}

	return nil
}
