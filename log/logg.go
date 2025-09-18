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
