package log

import (
	"dashboard-backend/pkg/setting"
	"os"
	"sort"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logChan = make(chan *logMsg, 10000)
)

type logMsg struct {
	level Level
	msg   *string
	time  time.Time
}

var zapLogger *zap.SugaredLogger

var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]*logMsg, 0, 1000)
	},
}

func InitZapLogger() {
	// 如果log不写入文件，不用处理日志
	if !setting.LogConfigSetting.LogOutFile && !setting.LogConfigSetting.LogOutConsole {
		return
	}
	encoder := getEncoder()

	cores := make([]zapcore.Core, 0)
	if setting.LogConfigSetting.LogOutFile {
		writeSyncer := getLogWriter()
		cores = append(cores, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
	}
	if setting.LogConfigSetting.LogOutConsole {
		cores = append(cores, zapcore.NewCore(encoder, os.Stdout, zapcore.DebugLevel))
	}

	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	zapLogger = logger.Sugar()

	// 启动异步日志workers
	go logWorker()
}

func logWorker() {
	defer close(logChan)

	// 从pool获取一个buffer
	buffer := bufferPool.Get().([]*logMsg)
	ticker := time.NewTicker(100 * time.Millisecond)

	// 创建输出channel
	outputChan := make(chan []*logMsg, 100)
	defer close(outputChan)

	// 启动输出goroutine
	go func() {
		for msgs := range outputChan {
			// 按时间排序
			sort.Slice(msgs, func(i, j int) bool {
				return msgs[i].time.Before(msgs[j].time)
			})
			// 批量写入日志
			for _, msg := range msgs {
				zapPrint(msg.level, *msg.msg)
			}
			// 使用完后将buffer放回pool
			//nolint:SA6002
			bufferPool.Put(msgs[:0])
		}
	}()

	for {
		select {
		case msg := <-logChan:
			buffer = append(buffer, msg)

		case <-ticker.C:
			if len(buffer) > 0 {
				// 发送当前buffer到输出channel
				outputChan <- buffer
				// 获取新的buffer
				buffer = bufferPool.Get().([]*logMsg)
			}
		}
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = ""
	encoderConfig.CallerKey = ""
	encoderConfig.LevelKey = ""
	encoderConfig.LineEnding = " "
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	logFileName := setting.LogConfigSetting.LogOutDir + setting.LogConfigSetting.LogOutFileName
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    setting.LogConfigSetting.MaxSize,
		MaxBackups: setting.LogConfigSetting.MaxBackups,
		MaxAge:     setting.LogConfigSetting.MaxAge,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func zapPrintWithTime(time time.Time, level Level, msg *string) {
	if setting.AppSetting.IsDev {
		// 开发环境下同步打日志
		zapPrint(level, *msg)
		return
	}
	if zapLogger == nil {
		return
	}
	logChan <- &logMsg{
		level: level,
		msg:   msg,
		time:  time,
	}
}

func zapPrint(level Level, msg string) {
	if zapLogger == nil {
		return
	}
	switch level {
	case DEBUG:
		zapLogger.Debug(msg)
	case INFO:
		zapLogger.Info(msg)
	case ERROR:
		zapLogger.Error(msg)
	case WARNING:
		zapLogger.Warn(msg)
	case FATAL:
		// 目前fatal也当error打，否则会结束进程
		zapLogger.Error(msg)
	}
}
