package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件写日志相关内容

// Logger 日志对象
type FileLogger struct {
	Level       LogLevel
	filePath    string //文件保存的路径
	fileName    string //文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// NewFileLoggger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxFileSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxFileSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl

}

// 初始化文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("OPEN LOG failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("OPEN err LOG failed,err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// 获取文件大小
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed:%v", err)
		return false
	}
	// 当前文件大小大于最大值，返回true
	return fileInfo.Size() >= f.maxFileSize

}

// 日志切割
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割
	nowStr := time.Now().Format("20060102150405000")

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed:%v\n", err)
		return nil, err
	}

	logName := path.Join(f.filePath, fileInfo.Name())       //拿到当前日志文件的完整路径
	newLogName := fmt.Sprintf("/%s.bak%s", logName, nowStr) //拼接一个备份文件的名字
	// 1.先关闭当前文件
	file.Close()
	// 2.备份一下rename
	os.Rename(logName, newLogName)
	// 3.打开一个新的文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed: %v\n", err)
		return nil, err
	}
	// 4.将打开的新日志文件对象赋值给f.fileObj
	return fileObj, nil
}

// 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// 往文件写入内容
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNumber := getInfo(3)

		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}

		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNumber, msg)
		// 如果要记录的日志大于等于error级别，还要再err日志文件中再记录一次
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNumber, msg)
		}
	}
}

// 判断是否需要记录该日志
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
