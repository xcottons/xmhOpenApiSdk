package xmhOpenApiSdk

import (
	"fmt"
	"io"
	"os"
)

const (
	LevelNull  Level = 0
	LevelError Level = 1
	LevelWarn  Level = 2
	LevelInfo  Level = 3
	LevelDebug Level = 4
)

var DefaultLeveledLogger LeveledLoggerInterface = &LeveledLogger{
	Level: LevelError,
}

type Level uint32

type LeveledLogger struct {
	Level          Level
	stderrOverride io.Writer
	stdoutOverride io.Writer
}

func (l *LeveledLogger) Debugf(format string, v ...interface{}) {
	if l.Level >= LevelDebug {
		fmt.Fprintf(l.stdout(), "[DEBUG] "+format+"\n", v...)
	}
}

func (l *LeveledLogger) Errorf(format string, v ...interface{}) {
	// Infof logs a debug message using Printf conventions.
	if l.Level >= LevelError {
		fmt.Fprintf(l.stderr(), "[ERROR] "+format+"\n", v...)
	}
}
func (l *LeveledLogger) Infof(format string, v ...interface{}) {
	if l.Level >= LevelInfo {
		fmt.Fprintf(l.stdout(), "[INFO] "+format+"\n", v...)
	}
}
func (l *LeveledLogger) Warnf(format string, v ...interface{}) {
	if l.Level >= LevelWarn {
		fmt.Fprintf(l.stderr(), "[WARN] "+format+"\n", v...)
	}
}
func (l *LeveledLogger) stderr() io.Writer {
	if l.stderrOverride != nil {
		return l.stderrOverride
	}

	return os.Stderr
}
func (l *LeveledLogger) stdout() io.Writer {
	if l.stdoutOverride != nil {
		return l.stdoutOverride
	}

	return os.Stdout
}

// LeveledLoggerInterface If you have your own logger, then implement this interface and replace DefaultLeveledLogger
type LeveledLoggerInterface interface {
	Debugf(format string, v ...interface{})

	Errorf(format string, v ...interface{})

	Infof(format string, v ...interface{})

	Warnf(format string, v ...interface{})
}
