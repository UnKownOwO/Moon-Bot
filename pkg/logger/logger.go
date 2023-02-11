package logger

import (
	"fmt"
	"log"
)

// LogLevel 输出等级
type LogLevel uint8

const (
	Loglevel_Debug LogLevel = iota // 调试
	LogLevel_Info                  // 正常
	LogLevel_Warn                  // 警告
	LogLevel_Error                 // 错误
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
	logInfo := NewLogInfo(Loglevel_Debug, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func Info(msg string, param ...any) {
	logInfo := NewLogInfo(LogLevel_Info, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func Warn(msg string, param ...any) {
	logInfo := NewLogInfo(LogLevel_Warn, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}

func Error(msg string, param ...any) {
	logInfo := NewLogInfo(LogLevel_Error, fmt.Sprintf(msg, param...))
	logger.logChan <- logInfo
}
