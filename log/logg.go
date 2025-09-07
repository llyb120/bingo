package log

import (
	"fmt"
	"github.com/llyb120/bingo/core"
	"log"
)

func Info(context core.Context, msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(msg, args...))
	//zapPrint(INFO, msg)
}

func Error(context core.Context, msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(msg, args...))
	//zapPrint(ERROR, msg)
}

func Debug(context core.Context, msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(msg, args...))
	//zapPrint(DEBUG, msg)
}

func Warn(context core.Context, msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(msg, args...))
	//zapPrint(WARN, msg)
}
