package golog

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

// BaseLogger formats and outputs log messages. It implements some of the basic
// functionality that all loggers must provide.
type BaseLogger struct {
	// name is the logger identifier name
	name string

	// level is the minimum level that will be outputted
	level int

	// tmpl is the template used to format log messages
	tmpl *template.Template
}

// NewBaseLogger creates a new BaseLogger instance
func NewBaseLogger(name string) *BaseLogger {
	logger := &BaseLogger{}

	logger.SetName(name)
	logger.SetLevel(DebugLevel)
	logger.SetFormatTmpl(DefaultLogFmt)

	return logger
}

func (l *BaseLogger) SetName(name string) {
	l.name = name
}

func (l *BaseLogger) SetLevel(level int) {
	l.level = level
}

func (l *BaseLogger) SetFormatTmpl(tmpl string) {
	l.tmpl = MustMkTmpl(tmpl)
}

// output writes a log message for a level to a writer
func (l BaseLogger) output(w io.Writer, level int, data ...interface{}) {
	// If logging to level which should not be displayed exit fn
	if level < l.level {
		return
	}

	// Get log level name from priority integer
	levelName := Levels[level]

	// Convert data to string
	strData := []string{}
	for _, d := range data {
		str := d.(string)
		strData = append(strData, str)
	}

	msg := strings.Join(strData, " ")

	// Format log message with template
	logCtx := NewLogMsgCtx(l.name, levelName, msg)

	s := MustExecTmpl(l.tmpl, logCtx)

	// Output
	fmt.Fprint(w, s)
}

// outputf formats a message and passes it to output
func (l BaseLogger) outputf(w io.Writer, level int, format string,
	data ...interface{}) {

	// If logging to level which should not be displayed exit fn
	if level < l.level {
		return
	}

	// Format message and pass to output fn
	msg := fmt.Sprintf(format, data...)

	l.output(w, level, msg)
}
