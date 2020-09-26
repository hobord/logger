package lgrs

import (
	"fmt"
	"strings"

	"github.com/hobord/logger"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (l *lgrs) WithError(err error) logger.Logger {
	errorEnty := logrusErrorEntry{
		Entry: l.base,
	}

	entry := errorEnty.WithStack(err)
	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}

type logrusErrorEntry struct {
	*logrus.Entry

	// Depth defines how much of the stacktrace you want.
	Depth int
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func (e *logrusErrorEntry) WithStack(err error) *logrus.Entry {
	out := e.Entry

	common := func(pError stackTracer) {
		st := pError.StackTrace()
		depth := 3
		if e.Depth != 0 {
			depth = e.Depth
		}
		if depth > len(st) {
			depth = len(st)
		}
		valued := fmt.Sprintf("%+v", st[0:depth])
		valued = strings.Replace(valued, "\t", "", -1)
		stack := strings.Split(valued, "\n")
		out = out.WithField("stack", stack[2:])
	}

	if err2, ok := err.(stackTracer); ok {
		common(err2)
	}

	if err2, ok := errors.Cause(err).(stackTracer); ok {
		common(err2)
	}

	return out.WithError(err)
}
