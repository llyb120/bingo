package log

// import (
// 	"bytes"
// 	"dashboard-backend/pkg/setting"
// 	"fmt"
// 	"log"
// 	"os"
// 	"regexp"
// 	"runtime"
// 	"strings"
// 	"time"

// 	"github.com/bytedance/sonic"
// 	"github.com/gin-gonic/gin"
// 	jsoniter "github.com/json-iterator/go"
// )

// type Level int

// var json = jsoniter.ConfigCompatibleWithStandardLibrary

// var (
// 	F                  *os.File
// 	today              string
// 	DefaultPrefix      = ""
// 	DefaultCallerDepth = 2

// 	logger        *log.Logger
// 	logPrefix     = ""
// 	levelFlags    = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
// 	replaceStrArr = []string{"\\n", "\\t"}

// 	// 仿slf4j的增强
// 	slf4jLogRegExp = regexp.MustCompile(`\{\}`)
// )

// const (
// 	DEBUG Level = iota
// 	INFO
// 	WARNING
// 	ERROR
// 	FATAL
// )

// type logContent struct {
// 	Date      string `json:"date"`
// 	Level     string `json:"level"`
// 	RequestId string `json:"requestId"`
// 	TraceId   string `json:"traceId"`
// 	UserName  string `json:"userName"`
// 	UserId    string `json:"userId"`
// 	File      string `json:"file"`
// 	Time      int    `json:"time"` // ms
// 	Msg       string `json:"msg"`
// }

// type statsContent struct {
// 	Date        string  `json:"date"`
// 	Level       string  `json:"level"`
// 	RequestId   string  `json:"requestId"`
// 	TraceId     string  `json:"traceId"`
// 	UserName    string  `json:"userName"`
// 	UserId      string  `json:"userId"`
// 	File        string  `json:"file"`
// 	Api         string  `json:"api"`
// 	Cache       int     `json:"cache"`
// 	SessionId   string  `json:"sessionId"`
// 	Status      int     `json:"status"`
// 	Time        float64 `json:"time"`
// 	System      string  `json:"system"`
// 	RequestPath string  `json:"request.path"`
// 	UserOpenid  string  `json:"user.openid"`
// 	Latency     float64 `json:"latency"`
// 	Ts          string  `json:"ts"`
// }

// type subApiStatsLog struct {
// 	Date      string  `json:"date"`
// 	Level     string  `json:"level"`
// 	RequestId string  `json:"requestId"`
// 	TraceId   string  `json:"traceId"`
// 	UserName  string  `json:"userName"`
// 	UserId    string  `json:"userId"`
// 	File      string  `json:"file"`
// 	Api       string  `json:"api"`
// 	SubApi    string  `json:"subApi"`
// 	Cache     int     `json:"cache"`
// 	SessionId string  `json:"sessionId"`
// 	Status    int     `json:"status"`
// 	Time      float64 `json:"time"`
// 	Msg       string  `json:"msg"`
// }

// // Debug output logs at debug level
// func Debug(c *gin.Context, v ...interface{}) {
// 	setLogContnt(c, DEBUG, fmt.Sprint(v...))
// }
// func DebugF(c *gin.Context, format string, v ...interface{}) {
// 	setLogContnt(c, DEBUG, fmt.Sprintf(format, v...))
// }

// // Info output logs at info level
// func Info(c *gin.Context, v ...interface{}) {
// 	setLogContnt(c, INFO, fmt.Sprint(v...))
// }
// func InfoF(c *gin.Context, format string, v ...interface{}) {
// 	setLogContnt(c, INFO, fmt.Sprintf(format, v...))
// }

// // InfoFF
// // 写一个增强，仿slf4j的方式，自动识别传入的参数
// // 使用{}进行占位符，使用者无需关注对应的参数类型
// // 使用方法 InfoFF(c, "test: {}", "test")
// func InfoFF(c *gin.Context, format string, v ...interface{}) {
// 	InfoF(c, handleFormat(format, v), v...)
// }

// // Warn output logs at warn level
// func Warn(c *gin.Context, v ...interface{}) {
// 	setLogContnt(c, WARNING, fmt.Sprint(v...))
// }

// // Warn output logs at warn level
// func WarnF(c *gin.Context, format string, v ...interface{}) {
// 	setLogContnt(c, WARNING, fmt.Sprintf(format, v...))
// }

// // Error output logs at error level
// func Error(c *gin.Context, v ...interface{}) {
// 	setLogContnt(c, ERROR, fmt.Sprint(v...))
// }
// func ErrorF(c *gin.Context, format string, v ...interface{}) {
// 	setLogContnt(c, ERROR, fmt.Sprintf(format, v...))
// }

// // Fatal output logs at fatal level
// func Fatal(c *gin.Context, v ...interface{}) {
// 	setLogContnt(c, FATAL, fmt.Sprint(v...))
// }
// func FatalF(c *gin.Context, format string, v ...interface{}) {
// 	setLogContnt(c, FATAL, fmt.Sprintf(format, v...))
// }

// func SkyWalkingF(c *gin.Context, t int, depth int, level Level, format string, v ...interface{}) {
// 	setLogContntSkyWaling(c, level, fmt.Sprintf(format, v...), depth, t)
// }

// const (
// 	Bold        = "\033[1m"
// 	ColorReset  = "\033[0m"
// 	ColorRed    = "\033[31m"
// 	ColorGreen  = "\033[32m"
// 	ColorYellow = "\033[33m"
// 	ColorOrange = "\033[38;5;208m"
// 	ColorPink   = "\033[38;5;213m"
// 	ColorBlue   = "\033[34m"
// 	ColorPurple = "\033[35m"
// 	ColorCyan   = "\033[36m"
// 	ColorWhite  = "\033[37m"
// 	ColorSky    = "\033[38;5;117m"
// )

// func setLogContnt(c *gin.Context, level Level, msg string) {
// 	// setPrefix set the prefix of the log output
// 	var (
// 		userId, userName, requestId, traceId string
// 	)

// 	p, file, line, _ := runtime.Caller(DefaultCallerDepth)

// 	if c != nil {
// 		requestId = c.GetString("Request-Id")
// 		traceId = c.GetString("Trace-Id")
// 		userContext, _ := c.Get("userinfo")

// 		if userContext != nil {
// 			userinfo := userContext.(map[string]interface{})
// 			userId = fmt.Sprintf("%s", userinfo["userid"])
// 			userName = fmt.Sprintf("%s", userinfo["username"])
// 		}
// 	}

// 	if setting.AppSetting.IsDev {
// 		DevLog(c, level, msg)
// 		return
// 	}

// 	logC := &logContent{
// 		Date:      now().Format("2006-01-02 15:04:05.000"),
// 		Level:     levelFlags[level],
// 		RequestId: requestId,
// 		TraceId:   traceId,
// 		UserName:  userName,
// 		UserId:    userId,
// 		File:      fmt.Sprintf("%s:%d %s", file, line, runtime.FuncForPC(p).Name()),
// 		Msg:       msg,
// 	}
// 	// log 中防止特殊字符被转义，影响log阅读
// 	bf := bytes.NewBuffer([]byte{})
// 	encoder := sonic.ConfigDefault.NewEncoder(bf)
// 	encoder.SetEscapeHTML(false)
// 	encoder.Encode(logC)
// 	logMsg := replace(bf.String())
// 	zapPrintWithTime(now(), level, &logMsg)
// 	return
// }

// func setLogContntSkyWaling(c *gin.Context, level Level, msg string, depth int, t int) {
// 	var (
// 		userId, userName, requestId, traceId string
// 	)

// 	p, file, line, _ := runtime.Caller(depth)

// 	if c != nil {
// 		requestId = c.GetString("Request-Id")
// 		traceId = c.GetString("Trace-Id")
// 		userContext, _ := c.Get("userinfo")
// 		if userContext != nil {
// 			userinfo := userContext.(map[string]interface{})
// 			userId = fmt.Sprintf("%v", userinfo["userid"])
// 			userName = fmt.Sprintf("%v", userinfo["username"])
// 		}
// 	}

// 	if setting.AppSetting.IsDev {
// 		DevLog(c, level, msg)
// 		return
// 	}

// 	logC := &logContent{
// 		Date:      now().Format("2006-01-02 15:04:05.000"),
// 		Level:     levelFlags[level],
// 		RequestId: requestId,
// 		TraceId:   traceId,
// 		UserName:  userName,
// 		UserId:    userId,
// 		File:      fmt.Sprintf("%s:%d %s", file, line, runtime.FuncForPC(p).Name()),
// 		Time:      t,
// 		Msg:       msg,
// 	}
// 	// log 中防止特殊字符被转义，影响log阅读
// 	bf := bytes.NewBuffer([]byte{})
// 	encoder := sonic.ConfigStd.NewEncoder(bf)
// 	encoder.SetEscapeHTML(false)
// 	encoder.Encode(logC)
// 	logMsg := replace(bf.String())
// 	zapPrintWithTime(now(), level, &logMsg)
// 	return
// }

// func replace(str string) string {
// 	for _, v := range replaceStrArr {
// 		str = strings.Replace(str, v, " ", -1)
// 	}
// 	return str
// }

// func GetStack() string {
// 	var buf [4096]byte
// 	n := runtime.Stack(buf[:], false)
// 	return string(buf[:n])
// }

// func StatsLog(c *gin.Context, api string, cache, status int, t float64) {
// 	var (
// 		userId, userName, requestId, traceId, sessionId string
// 	)

// 	p, file, line, _ := runtime.Caller(DefaultCallerDepth)

// 	if c != nil {
// 		requestId = c.GetString("Request-Id")
// 		traceId = c.GetString("Trace-Id")
// 		sessionId = c.GetHeader("Session-Id")
// 		userContext, _ := c.Get("userinfo")
// 		if userContext != nil {
// 			userinfo := userContext.(map[string]interface{})
// 			userId = fmt.Sprintf("%s", userinfo["userid"])
// 			userName = fmt.Sprintf("%s", userinfo["username"])
// 		}
// 	}

// 	if setting.AppSetting.IsDev {
// 		DevLog(c, INFO, fmt.Sprintf("api: %s, cache: %d, status: %d, time: %f", api, cache, status, t))
// 		return
// 	}
// 	date := now().Format("2006-01-02 15:04:05.000")
// 	logC := &statsContent{
// 		Date:        date,
// 		Api:         api,
// 		Cache:       cache,
// 		SessionId:   sessionId,
// 		Status:      status,
// 		Time:        t,
// 		Level:       levelFlags[INFO],
// 		RequestId:   requestId,
// 		TraceId:     traceId,
// 		UserName:    userName,
// 		UserId:      userId,
// 		File:        fmt.Sprintf("%s:%d %s", file, line, runtime.FuncForPC(p).Name()),
// 		System:      setting.LogConfigSetting.System,
// 		RequestPath: c.Request.URL.Path,
// 		UserOpenid:  userName,
// 		Latency:     t,
// 		Ts:          date,
// 	}
// 	// log 中防止特殊字符被转义，影响log阅读
// 	bf := bytes.NewBuffer([]byte{})
// 	encoder := sonic.ConfigStd.NewEncoder(bf)
// 	encoder.SetEscapeHTML(false)
// 	encoder.Encode(logC)
// 	logMsg := replace(bf.String())
// 	zapPrintWithTime(now(), INFO, &logMsg)
// 	return
// }

// func SubApiStatsLog(c *gin.Context, api, subApi string, cache, status int, t float64, msg string) {
// 	var (
// 		userId, userName, requestId, traceId, sessionId string
// 	)

// 	p, file, line, _ := runtime.Caller(7)

// 	if c != nil {
// 		requestId = c.GetString("Request-Id")
// 		traceId = c.GetString("Trace-Id")
// 		sessionId = c.GetHeader("Session-Id")
// 		userContext, _ := c.Get("userinfo")
// 		if userContext != nil {
// 			userinfo := userContext.(map[string]interface{})
// 			userId = fmt.Sprintf("%s", userinfo["userid"])
// 			userName = fmt.Sprintf("%s", userinfo["username"])
// 		}
// 	}

// 	if setting.AppSetting.IsDev {
// 		DevLog(c, INFO, fmt.Sprintf("api: %s, subApi: %s, cache: %d, status: %d, time: %f, msg: %s", api, subApi, cache, status, t, msg))
// 		return
// 	}

// 	logC := &subApiStatsLog{
// 		Date:      now().Format("2006-01-02 15:04:05.000"),
// 		Api:       api,
// 		SubApi:    subApi,
// 		Cache:     cache,
// 		SessionId: sessionId,
// 		Status:    status,
// 		Time:      t,
// 		Level:     levelFlags[INFO],
// 		RequestId: requestId,
// 		TraceId:   traceId,
// 		UserName:  userName,
// 		UserId:    userId,
// 		File:      fmt.Sprintf("%s:%d %s", file, line, runtime.FuncForPC(p).Name()),
// 		Msg:       msg,
// 	}
// 	// log 中防止特殊字符被转义，影响log阅读
// 	bf := bytes.NewBuffer([]byte{})
// 	encoder := sonic.ConfigStd.NewEncoder(bf)
// 	encoder.SetEscapeHTML(false)
// 	encoder.Encode(logC)
// 	logMsg := replace(bf.String())
// 	zapPrintWithTime(now(), INFO, &logMsg)
// 	return
// }

// func handleFormat(format string, v []interface{}) string {
// 	var i = 0
// 	var ln = len(v)
// 	format = slf4jLogRegExp.ReplaceAllStringFunc(format, func(s string) string {
// 		if i < ln {
// 			switch (v)[i].(type) {
// 			case string:
// 				return "%s"
// 			case int:
// 				return "%d"
// 			case float64:
// 				return "%f"
// 			default:
// 				return "%v"
// 			}
// 		}
// 		i++
// 		return s
// 	})
// 	return format
// }

// func DevLog(c *gin.Context, level Level, msg string) {
// 	if !setting.AppSetting.IsDev {
// 		return
// 	}
// 	var color string
// 	switch level {
// 	case ERROR:
// 		color = ColorRed
// 	default:
// 		color = ColorOrange
// 	}
// 	// sql 标记成蓝色
// 	sqlRegexp := regexp.MustCompile(`sql=(.*?) (,execResult|,error)`)
// 	msg = sqlRegexp.ReplaceAllStringFunc(msg, func(s string) string {
// 		subMatch := sqlRegexp.FindStringSubmatch(s)
// 		if len(subMatch) > 1 {
// 			return fmt.Sprintf("sql= %s%s%s%s %s", ColorBlue, subMatch[1], Bold, color, subMatch[2])
// 		}
// 		return s
// 	})
// 	msg = fmt.Sprintf("%s%s[%s] [%s] %s %s \n\n", Bold, color, now().Format("2006-01-02 15:04:05.000"), levelFlags[level], msg, ColorReset)
// 	zapPrintWithTime(now(), level, &msg)
// }

// var location *time.Location

// func init() {
// 	var err error
// 	location, err = time.LoadLocation("Asia/Shanghai")
// 	if err != nil {
// 		// If fails to load location, fallback to UTC time and log error.
// 		log.Println("Failed to load Asia/Shanghai location:", err)
// 	}
// }

// func now() time.Time {
// 	if location == nil {
// 		return time.Now()
// 	}
// 	return time.Now().In(location)
// }
