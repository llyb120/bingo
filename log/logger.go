package log

import (
	"fmt"

	"github.com/llyb120/bingo/core"
)

type Logger interface {
	Info(context core.Context, msg string, args ...interface{})
	Error(context core.Context, msg string, args ...interface{})
	Debug(context core.Context, msg string, args ...interface{})
	Warn(context core.Context, msg string, args ...interface{})
}

type defaultLogger struct {
}

func (d defaultLogger) Info(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("INFO", ColorPurple, msg)
}

func (d defaultLogger) Error(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("ERROR", ColorRed, msg)
}

func (d defaultLogger) Debug(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("DEBUG", ColorBlue, msg)
}

func (d defaultLogger) Warn(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("WARN", ColorYellow, msg)
}
