package logger

import (
	"context"
	"io"
)

// Fields type, used to pass to `WithFields`.
type Fields = map[string]interface{}

// Logger is the interface for loggers
type Logger interface {
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})

	Log(string, ...interface{})
	Logln(string, ...interface{})
	Logf(string, string, ...interface{})

	WithField(string, interface{}) Logger
	WithFields(Fields) Logger
	SetDefaultField(string, interface{})
	SetDefaultFields(Fields)

	WithError(error) Logger

	SetLevel(Level)
	SetOutput(output io.Writer)

	WithCorelationID(corelationID CorelationID) Logger
	SetDefaulthCorelationID(corelationID CorelationID)

	WithTraceID(traceID TraceID) Logger
	SetDefaultTraceID(traceID TraceID)

	WithContext(context.Context) Logger
	SetDefaultContext(ctx context.Context)

	WithTags([]string) Logger
	SetDefaultTags([]string)
	GetTags() []string
}
