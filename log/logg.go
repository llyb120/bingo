package log

import (
	"fmt"
	"time"

	"github.com/llyb120/bingo/core"
)

// ANSI 颜色代码定义
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[38;5;210m" // 胭脂红 - 如思君不见的红颜
	ColorGreen  = "\033[38;5;108m" // 青山绿 - 如远山含黛的青翠
	ColorYellow = "\033[38;5;222m" // 烟雨黄 - 如江南烟雨中的暖阳
	ColorBlue   = "\033[38;5;109m" // 天青色 - 如雨后初霁的青瓷
	ColorPurple = "\033[38;5;146m" // 丁香紫 - 如古巷中的幽香
	ColorCyan   = "\033[38;5;152m" // 烟波青 - 如水墨画中的远山
	ColorWhite  = "\033[38;5;255m" // 宣纸白 - 如诗笺上的纯净
	Bold        = "\033[1m"
)

var systemLogger = &defaultLogger{}
var loggerFactory = core.Use[Logger]()

// colorPrint 打印带颜色的日志
func colorPrint(level, color, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	coloredMsg := fmt.Sprintf("%s%s[%s] [%s] %s%s", Bold, color, timestamp, level, msg, ColorReset)
	fmt.Println(coloredMsg)
}

func Info(context core.Context, msg string, args ...interface{}) {
	//zapPrint(INFO, msg)
	logger := getLogger()
	if logger != nil {
		logger.Info(context, msg, args...)
	}
}

func Error(context core.Context, msg string, args ...interface{}) {
	//zapPrint(ERROR, msg)
	logger := getLogger()
	if logger != nil {
		logger.Error(context, msg, args...)
	}
}

func Debug(context core.Context, msg string, args ...interface{}) {
	logger := getLogger()
	if logger != nil {
		logger.Debug(context, msg, args...)
	}
}

func Warn(context core.Context, msg string, args ...interface{}) {
	logger := getLogger()
	if logger != nil {
		logger.Warn(context, msg, args...)
	}
}

func getLogger() Logger {
	logger := loggerFactory()
	if logger != nil {
		return logger
	}
	return systemLogger
}
