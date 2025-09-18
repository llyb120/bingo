package log

import (
	"fmt"
	"log"
	"time"

	"github.com/llyb120/bingo/core"
)

// ANSI 颜色代码定义
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	Bold        = "\033[1m"
)

// colorPrint 打印带颜色的日志
func colorPrint(level, color, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	coloredMsg := fmt.Sprintf("%s%s[%s] [%s] %s%s", Bold, color, timestamp, level, msg, ColorReset)
	log.Println(coloredMsg)
}

func Info(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("INFO", ColorGreen, msg)
	//zapPrint(INFO, msg)
}

func Error(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("ERROR", ColorRed, msg)
	//zapPrint(ERROR, msg)
}

func Debug(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("DEBUG", ColorBlue, msg)
	//zapPrint(DEBUG, msg)
}

func Warn(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("WARN", ColorYellow, msg)
	//zapPrint(WARN, msg)
}
