package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// Logger 日志对象
type Logger struct {
	Level LogLevel
}

// NewLog 构造Logger
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) log(lv LogLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNumber := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNumber, msg)
	}
}

func (l Logger) enable(level LogLevel) bool {
	return level >= l.Level
}

func (l Logger) Debug(format string, a ...interface{}) {

	l.log(DEBUG, format, a...)

}

func (l Logger) Info(format string, a ...interface{}) {

	l.log(INFO, format, a...)

}

func (l Logger) Warning(format string, a ...interface{}) {

	l.log(WARNING, format, a...)

}

func (l Logger) Error(format string, a ...interface{}) {

	l.log(ERROR, format, a...)

}
func (l Logger) Fatal(format string, a ...interface{}) {

	l.log(FATAL, format, a...)

}
