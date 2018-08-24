package golog

// DefaultLogFmt is the default log format template
const DefaultLogFmt string = "{{ .Name }} [{{ .Level }}] {{ .Msg }}\n"

// Logger defines methods for outputting information
type Logger interface {
	// SetName sets the logger's identifying name
	SetName(name string)

	// SetLevel sets the minimum log level which will be outputted
	SetLevel(level int)

	// SetFormatTmpl defines a Go template used to format log output.
	// Fields from the LogMsgCtx struct are available to use in the
	// template.
	SetFormatTmpl(tmpl string)

	// Fatal writes at the FATAL level and panics the process
	Fatal(data ...interface{})

	// Fatalf formats a string, writes at the FATAL level, and panics the
	// process
	Fatalf(format string, data ...interface{})

	// Error writes at the ERROR level
	Error(data ...interface{})

	// Errorf formats a string and writes at the ERROR level
	Errorf(format string, data ...interface{})

	// Warn writes at the WARN level
	Warn(data ...interface{})

	// Warnf formats a string and writes at the WARN level
	Warnf(format string, data ...interface{})

	// Info writes at the INFO level
	Info(data ...interface{})

	// Infof formats a string and writes at the INFO level
	Infof(format string, data ...interface{})

	// Debug writes at the DEBUG level
	Debug(data ...interface{})

	// Debugf formats a string and writes at the DEBUG level
	Debugf(format string, data ...interface{})
}
