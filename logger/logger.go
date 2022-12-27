package logger

import "fmt"

// Logger interface API for log.Logger.
type Logger interface {
	Printf(string, ...interface{})
}

type DefaultLogger struct{}

func (d *DefaultLogger) Printf(msg string, args ...interface{}) {
	fmt.Printf(msg, args)
}

var Log = DefaultLogger{}

type LoggerFunc func(string, ...interface{})

func (f LoggerFunc) Printf(msg string, args ...interface{}) { f(msg, args...) }
