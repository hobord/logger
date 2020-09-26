package lgrs

import (
	"context"

	"github.com/hobord/logger"
	"github.com/sirupsen/logrus"
)

func addCorelationIDToEntry(entry *logrus.Entry, corelationID logger.CorelationID) *logrus.Entry {
	return entry.WithField(logger.LogFieldKeyCorelationID, corelationID.String())
}
func (l *lgrs) WithCorelationID(corelationID logger.CorelationID) logger.Logger {
	entry := addCorelationIDToEntry(l.base, corelationID)
	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}
func (l *lgrs) SetDefaulthCorelationID(corelationID logger.CorelationID) {
	l.base = addCorelationIDToEntry(l.base, corelationID)
}

func addTraceIDToEntry(entry *logrus.Entry, traceID logger.TraceID) *logrus.Entry {
	return entry.WithFields(logrus.Fields{
		"TraceIDCurrent": traceID.GetCurrent(),
		"TraceIDPrev":    traceID.GetPrev(),
	})
}
func (l *lgrs) WithTraceID(traceID logger.TraceID) logger.Logger {
	entry := addTraceIDToEntry(l.base, traceID)
	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}
func (l *lgrs) SetDefaultTraceID(traceID logger.TraceID) {
	l.base = addTraceIDToEntry(l.base, traceID)
}

func (l *lgrs) withContext(ctx context.Context) *logrus.Entry {
	entry := l.base.WithContext(ctx)

	// if context has corelationID then add into the log
	corelationID := logger.GetCorelationIDFromContext(ctx)
	if corelationID != nil {
		entry = addCorelationIDToEntry(entry, *corelationID)
	}

	// if context has traceID then add into the log
	traceID := logger.GetTraceIDFromContext(ctx)
	if traceID != nil {
		entry = addTraceIDToEntry(entry, *traceID)
	}

	return entry
}
func (l *lgrs) WithContext(ctx context.Context) logger.Logger {
	entry := l.withContext(ctx)

	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}
func (l *lgrs) SetDefaultContext(ctx context.Context) {
	l.base = l.withContext(ctx)
}
