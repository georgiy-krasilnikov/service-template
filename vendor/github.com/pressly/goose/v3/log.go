package goose

import (
	std "log"
)

var log Logger = &stdLogger{}

// Logger is standard logger interface
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// SetLogger sets the logger for package output
func SetLogger(l Logger) {
	log = l
}

// stdLogger is a default logger that outputs to a stdlib's log.std logger.
type stdLogger struct{}

func (*stdLogger) Fatal(v ...interface{})                 { std.Fatal(v...) }
func (*stdLogger) Fatalf(format string, v ...interface{}) { std.Fatalf(format, v...) }
func (*stdLogger) Print(v ...interface{})                 { std.Print(v...) }
func (*stdLogger) Println(v ...interface{})               { std.Println(v...) }
func (*stdLogger) Printf(format string, v ...interface{}) { std.Printf(format, v...) }

// NopLogger returns a logger that discards all logged output.
func NopLogger() Logger {
	return &nopLogger{}
}

type nopLogger struct{}

var _ Logger = (*nopLogger)(nil)

func (*nopLogger) Fatal(v ...interface{})                 {}
func (*nopLogger) Fatalf(format string, v ...interface{}) {}
func (*nopLogger) Print(v ...interface{})                 {}
func (*nopLogger) Println(v ...interface{})               {}
func (*nopLogger) Printf(format string, v ...interface{}) {}
