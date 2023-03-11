package logger

import (
	"fmt"
	"log"
	"time"
)

// 控制台颜色
var (
	GreenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	WhiteBg   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	YellowBg  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	RedBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	BlueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	MagentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	CyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	Green     = string([]byte{27, 91, 51, 50, 109})
	White     = string([]byte{27, 91, 51, 55, 109})
	Yellow    = string([]byte{27, 91, 51, 51, 109})
	Red       = string([]byte{27, 91, 51, 49, 109})
	Blue      = string([]byte{27, 91, 51, 52, 109})
	Magenta   = string([]byte{27, 91, 51, 53, 109})
	Cyan      = string([]byte{27, 91, 51, 54, 109})
	Reset     = string([]byte{27, 91, 48, 109})
)

// LogLevel 输出等级
type LogLevel uint8

const (
	LoglevelDebug LogLevel = iota // 调试
	LogLevelInfo                  // 正常
	LogLevelWarn                  // 警告
	LogLevelError                 // 错误
)

type LogInfo struct {
	Level LogLevel // 输出等级
	Msg   string   // 格式化后的字符串
}

func NewLogInfo(level LogLevel, msg string) *LogInfo {
	logInfo := new(LogInfo)
	logInfo.Level = level
	logInfo.Msg = msg
	return logInfo
}

type Logger struct {
	logChan chan *LogInfo // 待处理消息通道
}

// logHandler 处理传入的消息
func (l *Logger) logHandler() {
	for {
		logInfo := <-l.logChan
		var logStr string
		nowTime := time.Now().Format("2006-01-02 15:04:05.000")
		// 输出格式
		switch logInfo.Level {
		case LoglevelDebug:
			logStr = fmt.Sprintf("%v%v %v[%v]:%v %v", White, nowTime, Blue, l.getLevelStr(logInfo.Level), Reset, logInfo.Msg)
		case LogLevelInfo:
			logStr = fmt.Sprintf("%v%v %v[%v]:%v %v", White, nowTime, Green, l.getLevelStr(logInfo.Level), Reset, logInfo.Msg)
		case LogLevelWarn:
			logStr = fmt.Sprintf("%v%v %v[%v]: %v%v", White, nowTime, Yellow, l.getLevelStr(logInfo.Level), logInfo.Msg, Reset)
		case LogLevelError:
			logStr = fmt.Sprintf("%v%v %v[%v]: %v%v", White, nowTime, Red, l.getLevelStr(logInfo.Level), logInfo.Msg, Reset)
		}
		log.Println(logStr)
	}
}

var logger *Logger

// InitLogger 初始化Logger
func InitLogger() {
	log.SetFlags(0)
	logger = new(Logger)
	logger.logChan = make(chan *LogInfo, 1000)
	go logger.logHandler()
}

func Debug(msg string, param ...any) {
	logInfo := NewLogInfo(LoglevelDebug, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func Info(msg string, param ...any) {
	logInfo := NewLogInfo(LogLevelInfo, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func Warn(msg string, param ...any) {
	logInfo := NewLogInfo(LogLevelWarn, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func Error(msg string, param ...any) {
	logInfo := NewLogInfo(LogLevelError, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func (l *Logger) getLevelInt(level string) LogLevel {
	switch level {
	case "DEBUG":
		return LoglevelDebug
	case "INFO":
		return LogLevelInfo
	case "WARN":
		return LogLevelWarn
	case "ERROR":
		return LogLevelError
	default:
		return LoglevelDebug
	}
}

func (l *Logger) getLevelStr(level LogLevel) string {
	switch level {
	case LoglevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return "DEBUG"
	}
}
