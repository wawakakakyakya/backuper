package logger

type actWriter interface {
	_write(msg string)
}

type LoggerInterface interface {
	write(ll loglevel, msg string, aw actWriter)
	Info(msg string)
	Error(err error)
	ErrorS(err string)
}
