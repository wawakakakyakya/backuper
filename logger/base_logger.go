package logger

import (
	"fmt"
	"time"
)

type loglevel string

var (
	Info  loglevel = "Info"
	Error loglevel = "Error"
)

type BaseLogger struct {
	LoggerInterface
}

func (l *BaseLogger) getTime() string {
	return time.Now().Format("2006/1/2 15:04:05")
}

func (l *BaseLogger) write(ll loglevel, msg string, aw actWriter) {
	now := l.getTime()
	aw._write(fmt.Sprintf("[%s][%s] %s", ll, now, msg))
}

func (l *BaseLogger) _info(msg string, aw actWriter) {
	l.write(Info, msg, aw)
}

func (l *BaseLogger) _errorS(err string, aw actWriter) {
	l.write(Error, err, aw)
}

func (l *BaseLogger) _error(err error, aw actWriter) {
	l.write(Error, err.Error(), aw)
}

// func (l *StdOutLogger) Error(err interface{}) {
// 	switch err := err.(type) {
// 	case string:
// 		l.write(Error, err)
// 	case error:
// 		l.write(Error, err.Error())
// 	default:
// 		panic("Invalid type!")
// 	}
// }

// var Logger StdOutLogger
