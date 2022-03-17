package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 日志等级
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota //0
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// 将传入日志等级转换为logLevel方便比较
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func getInfo(n int) (funcName, fileName string, lineNumber int) {
	// ok:能否取到，file：拿到调用该函数的文件，line：第多少行调用,n是函数调用在第几层
	pc, file, lineNumber, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNumber
}
