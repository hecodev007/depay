package loger

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("0102-150405.000")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf("%s [%s:%d][%s] %s\n", timestamp, file, len, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

type logFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
	encoding string
}
//var FileDate = "20220213"
func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}

	//判断是否需要切换日期
	fileDate := time.Now().Format("20060102")
	//fileDate := FileDate
	if p.fileDate != fileDate {
		p.file.Close()
		err = os.MkdirAll(fmt.Sprintf("%s/%s", p.logPath, fileDate), os.ModePerm)
		if err != nil {
			return 0, err
		}
		filename := fmt.Sprintf("%s/%s/%s-%s.log", p.logPath, fileDate, p.appName, fileDate)

		p.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println("log write err:", err)
			return 0, err
		}
		//fileWriter := logFileWriter{p.file, p.logPath, fileDate, p.appName, p.encoding}
		//log.SetOutput(&fileWriter)
		//log.SetReportCaller(true)
		//log.SetFormatter(new(LogFormatter))
	}
	//if p.encoding != "" {
	//	dataToEncode := ConvertStringToByte(string(data), p.encoding)
	//	n, e := p.file.Write(dataToEncode)
	//	return n, e
	//}

	n, e := p.file.Write(data)
	return n, e

}

func InitLog(logPath, appName, encoding string) {

	fileDate := time.Now().Format("20060102")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		log.Error(err)
		return
	}

	filename := fmt.Sprintf("%s/%s/%s-%s.log", logPath, fileDate, appName, fileDate)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Error(err)
		return
	}

	fileWriter := logFileWriter{file, logPath, fileDate, appName, encoding}
	log.SetOutput(&fileWriter)

	log.SetReportCaller(true)
	log.SetFormatter(new(LogFormatter))

}
