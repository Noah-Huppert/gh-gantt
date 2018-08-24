package golog

// LogMsgCtx holds information about a log message
type LogMsgCtx struct {
	// Name is the identifier of the logger
	Name string

	// Level is the name of the message log level
	Level string

	// Msg is the log message contents
	Msg string
}

// NewLogMsgCtx creates a new LogMsgCtx
func NewLogMsgCtx(name, lvl, msg string) LogMsgCtx {
	return LogMsgCtx{
		Name:  name,
		Level: lvl,
		Msg:   msg,
	}
}
