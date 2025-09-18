package log

import (
	"fmt"
	"time"

	"github.com/llyb120/bingo/core"
)

// ANSI 颜色代码定义
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[38;5;203m" // 柔和的红色
	ColorGreen  = "\033[38;5;77m"  // 柔和的绿色
	ColorYellow = "\033[38;5;185m" // 柔和的黄色
	ColorBlue   = "\033[38;5;75m"  // 柔和的蓝色
	ColorPurple = "\033[38;5;141m" // 柔和的紫色
	ColorCyan   = "\033[38;5;86m"  // 柔和的青色
	ColorWhite  = "\033[38;5;255m" // 柔和的白色
	Bold        = "\033[1m"
)

// colorPrint 打印带颜色的日志
func colorPrint(level, color, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	coloredMsg := fmt.Sprintf("%s%s[%s] [%s] %s%s", Bold, color, timestamp, level, msg, ColorReset)
	fmt.Println(coloredMsg)
}

func Info(context core.Context, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	colorPrint("INFO", ColorPurple, msg)
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
