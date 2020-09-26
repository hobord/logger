package lgrs

import (
	"io"
	"os"

	"github.com/hobord/logger"
	"github.com/sirupsen/logrus"
)

type lgrs struct {
	logger *logrus.Logger
	base   *logrus.Entry
}

// MakeLoggerWithLogrus is create a Logger wrapp logrus
func MakeLoggerWithLogrus(baseLogger *logrus.Logger) *lgrs {
	if baseLogger == nil {
		baseLogger = logrus.New()
		baseLogger.Out = os.Stdout
		baseLogger.SetFormatter(&logrus.JSONFormatter{})
	}
	entry := logrus.NewEntry(baseLogger)
	return &lgrs{
		logger: baseLogger,
		base:   entry,
	}
}

func (l *lgrs) SetLevel(level logger.Level) {
	logrusLevel := l.convertLevel(level)
	l.logger.SetLevel(logrusLevel)
}

func (l *lgrs) SetOutput(output io.Writer) {
	l.logger.SetOutput(output)
}

func (l *lgrs) cloneLgrs() *lgrs {
	return &lgrs{
		logger: l.logger,
		base:   l.base,
	}
}

func (l *lgrs) convertLevel(level logger.Level) logrus.Level {
	switch level {
	case logger.PanicLevel:
		return logrus.PanicLevel
	case logger.FatalLevel:
		return logrus.FatalLevel
	case logger.ErrorLevel:
		return logrus.ErrorLevel
	case logger.WarnLevel:
		return logrus.WarnLevel
	case logger.InfoLevel:
		return logrus.InfoLevel
	case logger.DebugLevel:
		return logrus.DebugLevel
	case logger.TraceLevel:
		return logrus.TraceLevel
	}
	return logrus.InfoLevel
}
