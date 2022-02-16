package logger

import "fmt"

type StdOutLogger struct {
	actWriter
	BaseLogger // call xxx instead of BaseLogger.xxx
}

func (l *StdOutLogger) _write(msg string) {
	fmt.Println(msg)
}

func (l *StdOutLogger) Info(msg string) {
	l._info(msg, l)
}

func (l *StdOutLogger) ErrorS(err string) {
	l._errorS(err, l)
}

func (l *StdOutLogger) Error(err error) {
	l._error(err, l)
}

func NewStdoutLogger() *StdOutLogger {
	return new(StdOutLogger)
}
