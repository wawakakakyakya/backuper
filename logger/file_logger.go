package logger

import "fmt"

type FileLogger struct {
	BaseLogger // call xxx instead of BaseLogger.xxx
}

func (f *FileLogger) _write(msg string) {
	fmt.Println(msg)
}

func (f *FileLogger) Info(msg string) {
	f._info(msg, f)
}

func (f *FileLogger) ErrorS(err string) {
	f._errorS(err, f)
}

func (f *FileLogger) Error(err error) {
	f._error(err, f)
}

func NewFileLogger() *FileLogger {
	return new(FileLogger)
}
