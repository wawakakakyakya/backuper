package logger

func NewLogger() LoggerInterface {
	return NewStdoutLogger()
}
