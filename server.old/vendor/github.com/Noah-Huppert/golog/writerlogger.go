package golog

import (
	"io"
	"os"
)

// WriterLogger implements the Logger interface using different io.Writers for
// the different log levels.
type WriterLogger struct {
	// BaseLogger is used to provide the basic functionality of a Logger
	*BaseLogger

	// FatalWriter is used to output FATAL level log messages
	FatalWriter io.Writer

	// ErrorWriter is used to output ERROR level log messages
	ErrorWriter io.Writer

	// WarnWriter is used to output WARN level log messages
	WarnWriter io.Writer

	// InfoWriter is used to output INFO level log messages
	InfoWriter io.Writer

	// DebugWriter is used to output DEBUG level log messages
	DebugWriter io.Writer
}

// NewStdLogger creates a WriterLogger which outputs normal messages to
// stdout and error messages to stderr
func NewStdLogger(name string) *WriterLogger {
	return &WriterLogger{
		BaseLogger:  NewBaseLogger(name),
		FatalWriter: os.Stderr,
		ErrorWriter: os.Stderr,
		WarnWriter:  os.Stderr,
		InfoWriter:  os.Stdout,
		DebugWriter: os.Stdout,
	}
}

// NewWriterLogger creates a new WriterLogger
func NewWriterLogger(name string, fatalWriter io.Writer, errorWriter io.Writer,
	warnWriter io.Writer, infoWriter io.Writer,
	debugWriter io.Writer) *WriterLogger {

	return &WriterLogger{
		BaseLogger:  NewBaseLogger(name),
		FatalWriter: fatalWriter,
		ErrorWriter: errorWriter,
		WarnWriter:  warnWriter,
		InfoWriter:  infoWriter,
		DebugWriter: debugWriter,
	}
}

func (l WriterLogger) Fatal(data ...interface{}) {
	l.output(l.FatalWriter, FatalLevel, data...)
	os.Exit(1)
}

func (l WriterLogger) Fatalf(format string, data ...interface{}) {
	l.outputf(l.FatalWriter, FatalLevel, format, data...)
	os.Exit(1)
}

func (l WriterLogger) Error(data ...interface{}) {
	l.output(l.ErrorWriter, ErrorLevel, data...)
}

func (l WriterLogger) Errorf(format string, data ...interface{}) {
	l.outputf(l.ErrorWriter, ErrorLevel, format, data...)
}

func (l WriterLogger) Warn(data ...interface{}) {
	l.output(l.WarnWriter, WarnLevel, data...)
}

func (l WriterLogger) Warnf(format string, data ...interface{}) {
	l.outputf(l.WarnWriter, WarnLevel, format, data...)
}

func (l WriterLogger) Info(data ...interface{}) {
	l.output(l.InfoWriter, InfoLevel, data...)
}

func (l WriterLogger) Infof(format string, data ...interface{}) {
	l.outputf(l.InfoWriter, InfoLevel, format, data...)
}

func (l WriterLogger) Debug(data ...interface{}) {
	l.output(l.DebugWriter, DebugLevel, data...)
}

func (l WriterLogger) Debugf(format string, data ...interface{}) {
	l.outputf(l.DebugWriter, DebugLevel, format, data...)
}
