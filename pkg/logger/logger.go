package logger

import (
	"fmt"
	"log"
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
		log.Println(logInfo.Msg)
	}
}

var logger *Logger

// InitLogger 初始化Logger
func InitLogger() {
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
