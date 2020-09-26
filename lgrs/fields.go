package lgrs

import (
	"github.com/hobord/logger"
	"github.com/sirupsen/logrus"
)

func (l *lgrs) WithField(key string, value interface{}) logger.Logger {
	entry := l.base.WithField(key, value)

	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}

func (l *lgrs) WithFields(fields logger.Fields) logger.Logger {
	lgFields := make(logrus.Fields, len(fields))
	for k, v := range fields {
		lgFields[k] = v
	}

	entry := l.base.WithFields(lgFields)

	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}

func (l *lgrs) SetDefaultField(key string, value interface{}) {
	l.base = l.base.WithField(key, value)
}

func (l *lgrs) SetDefaultFields(fields logger.Fields) {
	l.base = l.base.WithFields(fields)
}
