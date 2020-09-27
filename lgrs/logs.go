package lgrs

import "github.com/hobord/logger"

func (l *lgrs) Trace(args ...interface{}) {
	l.base.Trace(args...)
}
func (l *lgrs) Traceln(args ...interface{}) {
	l.base.Traceln(args...)
}
func (l *lgrs) Tracef(format string, args ...interface{}) {
	l.base.Tracef(format, args...)
}

func (l *lgrs) Debug(args ...interface{}) {
	l.base.Debug(args...)
}
func (l *lgrs) Debugln(args ...interface{}) {
	l.base.Debugln(args...)
}
func (l *lgrs) Debugf(format string, args ...interface{}) {
	l.base.Debugf(format, args...)
}

func (l *lgrs) Info(args ...interface{}) {
	l.base.Info(args...)
}
func (l *lgrs) Infoln(args ...interface{}) {
	l.base.Infoln(args...)
}
func (l *lgrs) Infof(format string, args ...interface{}) {
	l.base.Infof(format, args...)
}

func (l *lgrs) Warn(args ...interface{}) {
	l.base.Warn(args...)
}
func (l *lgrs) Warnln(args ...interface{}) {
	l.base.Warnln(args...)
}
func (l *lgrs) Warnf(format string, args ...interface{}) {
	l.base.Warnf(format, args...)
}

func (l *lgrs) Error(args ...interface{}) {
	l.base.Error(args...)
}
func (l *lgrs) Errorln(args ...interface{}) {
	l.base.Errorln(args...)
}
func (l *lgrs) Errorf(format string, args ...interface{}) {
	l.base.Errorf(format, args...)
}

func (l *lgrs) Fatal(args ...interface{}) {
	l.base.Fatal(args...)
}
func (l *lgrs) Fatalln(args ...interface{}) {
	l.base.Fatalln(args...)
}
func (l *lgrs) Fatalf(format string, args ...interface{}) {
	l.base.Fatalf(format, args...)
}

func (l *lgrs) Panic(args ...interface{}) {
	l.base.Panic(args...)
}
func (l *lgrs) Panicln(args ...interface{}) {
	l.base.Panicln(args...)
}
func (l *lgrs) Panicf(format string, args ...interface{}) {
	l.base.Panicf(format, args...)
}

func (l *lgrs) Log(level string, args ...interface{}) {
	lev, err := logger.ParseLevel(level)
	if err != nil {
		return
	}
	logrusLevel := l.convertLevel(lev)
	l.base.Log(logrusLevel, args...)
}
func (l *lgrs) Logln(level string, args ...interface{}) {
	lev, err := logger.ParseLevel(level)
	if err != nil {
		return
	}
	logrusLevel := l.convertLevel(lev)
	l.base.Logln(logrusLevel, args...)
}
func (l *lgrs) Logf(level string, format string, args ...interface{}) {
	lev, err := logger.ParseLevel(level)
	if err != nil {
		return
	}
	logrusLevel := l.convertLevel(lev)
	l.base.Logf(logrusLevel, format, args...)
}
