package out

import (
	"fmt"
	"strings"
)

type LogLevel int8

const (
	DEBUG LogLevel = iota
	INFO  LogLevel = iota
	WARN  LogLevel = iota
	ERROR LogLevel = iota
)

type Logger struct {
	logLevel LogLevel
}

func NewLogger(level string) Logger {

	var logLevel LogLevel
	switch strings.ToLower(level) {
	case "info":
		logLevel = INFO
	case "warn":
		logLevel = WARN
	case "error":
		logLevel = ERROR
	default:
		logLevel = DEBUG
	}

	return Logger{logLevel}
}

func (l Logger) Debug(args ...any) {
	if l.logLevel >= DEBUG {
		fmt.Println(args...)
	}
}
func (l Logger) Info(args ...any) {
	if l.logLevel >= INFO {
		fmt.Println(args...)
	}
}

func (l Logger) Warn(args ...any) {
	if l.logLevel >= WARN {
		fmt.Println(args...)
	}
}

func (l Logger) Error(args ...any) {
	if l.logLevel >= ERROR {
		fmt.Println(args...)
	}
}

func (l Logger) Log(level LogLevel, args ...any) {
	if l.logLevel >= level {
		fmt.Println(args...)
	}
}
