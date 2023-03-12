package logger

import (
	"bytes"
	"fmt"
	"log"
	"path"
	"runtime"
	"strings"
	"time"
)

// 控制台颜色
var (
	greenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green     = string([]byte{27, 91, 51, 50, 109})
	white     = string([]byte{27, 91, 51, 55, 109})
	yellow    = string([]byte{27, 91, 51, 51, 109})
	red       = string([]byte{27, 91, 51, 49, 109})
	blue      = string([]byte{27, 91, 51, 52, 109})
	magenta   = string([]byte{27, 91, 51, 53, 109})
	cyan      = string([]byte{27, 91, 51, 54, 109})
	reset     = string([]byte{27, 91, 48, 109})
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
	Level       LogLevel // 输出等级
	Msg         string   // 格式化后的字符串
	FileName    string   // 文件名
	FuncName    string   // 函数名
	Line        int      // 行号
	GoroutineId string   // 协程Id
	ThreadId    string   // 线程Id
}

type Logger struct {
	logChan chan *LogInfo // 待处理消息通道
	Level   LogLevel      // 最小输出等级
	Track   bool          // 是否详细输出
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
			logStr = fmt.Sprintf("%v%v %v[%v]:%v %v", white, nowTime, blue, l.getLevelStr(logInfo.Level), reset, logInfo.Msg)
		case LogLevelInfo:
			logStr = fmt.Sprintf("%v%v %v[%v]:%v %v", white, nowTime, green, l.getLevelStr(logInfo.Level), reset, logInfo.Msg)
		case LogLevelWarn:
			logStr = fmt.Sprintf("%v%v %v[%v]: %v%v", white, nowTime, yellow, l.getLevelStr(logInfo.Level), logInfo.Msg, reset)
		case LogLevelError:
			logStr = fmt.Sprintf("%v%v %v[%v]: %v%v", white, nowTime, red, l.getLevelStr(logInfo.Level), logInfo.Msg, reset)
		}
		if l.Track {
			logStr += fmt.Sprintf(" %v[%v:%v %v() goroutine:%v thread:%v]%v", magenta, logInfo.FileName, logInfo.Line, logInfo.FuncName, logInfo.GoroutineId, logInfo.ThreadId, reset)
		}
		logStr += "\n"
		log.Print(logStr)
	}
}

var logger *Logger

// InitLogger 初始化Logger
func InitLogger() {
	log.SetFlags(0)
	logger = new(Logger)
	logger.logChan = make(chan *LogInfo, 1000)
	logger.Track = true

	go logger.logHandler()
}

func Debug(msg string, param ...any) {
	if logger.Level > LoglevelDebug {
		return
	}
	logInfo := new(LogInfo)
	logInfo.Level = LoglevelDebug
	logInfo.Msg = fmt.Sprintf(msg, param...)
	if logger.Track {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = logger.getLineFunc()
		logInfo.GoroutineId = logger.getGoroutineId()
		logInfo.ThreadId = logger.getThreadId()
	}
	logger.logChan <- logInfo
}

func Info(msg string, param ...any) {
	if logger.Level > LogLevelInfo {
		return
	}
	logInfo := new(LogInfo)
	logInfo.Level = LogLevelInfo
	logInfo.Msg = fmt.Sprintf(msg, param...)
	if logger.Track {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = logger.getLineFunc()
		logInfo.GoroutineId = logger.getGoroutineId()
		logInfo.ThreadId = logger.getThreadId()
	}
	logger.logChan <- logInfo
}

func Warn(msg string, param ...any) {
	if logger.Level > LogLevelWarn {
		return
	}
	logInfo := new(LogInfo)
	logInfo.Level = LogLevelWarn
	logInfo.Msg = fmt.Sprintf(msg, param...)
	if logger.Track {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = logger.getLineFunc()
		logInfo.GoroutineId = logger.getGoroutineId()
		logInfo.ThreadId = logger.getThreadId()
	}
	logger.logChan <- logInfo
}

func Error(msg string, param ...any) {
	if logger.Level > LogLevelError {
		return
	}
	logInfo := new(LogInfo)
	logInfo.Level = LogLevelError
	logInfo.Msg = fmt.Sprintf(msg, param...)
	if logger.Track {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = logger.getLineFunc()
		logInfo.GoroutineId = logger.getGoroutineId()
		logInfo.ThreadId = logger.getThreadId()
	}
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

func (l *Logger) getGoroutineId() (goroutineId string) {
	buf := make([]byte, 32)
	runtime.Stack(buf, false)
	buf = bytes.TrimPrefix(buf, []byte("goroutine "))
	buf = buf[:bytes.IndexByte(buf, ' ')]
	goroutineId = string(buf)
	return goroutineId
}

func (l *Logger) getLineFunc() (fileName string, line int, funcName string) {
	var pc uintptr
	var file string
	var ok bool
	pc, file, line, ok = runtime.Caller(2)
	if !ok {
		return "???", -1, "???"
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	split := strings.Split(funcName, ".")
	if len(split) != 0 {
		funcName = split[len(split)-1]
	}
	return fileName, line, funcName
}

func Stack() string {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
}

func StackAll() string {
	buf := make([]byte, 1024*16)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
}
