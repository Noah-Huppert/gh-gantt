package golog

// FatalLevelName is the name of the FATAL log level
const FatalLevelName string = "FATAL"

// FatalLevel is used to identify the FATAL log level and compare it to
// other levels
const FatalLevel int = 100

// ErrorLevelName is the name of the ERROR log level
const ErrorLevelName string = "ERROR"

// ErrorLevel is used to identify the ERROR log level and compare it to
// other levels
const ErrorLevel int = 90

// WarnLevelName is the name of the WARN log level
const WarnLevelName string = "WARN"

// WarnLevel is used to identify the WARN log level and compare it to
// other levels
const WarnLevel int = 80

// InfoLevelName is the name of the INFO log level
const InfoLevelName string = "INFO"

// InfoLevel is used to identify the INFO log level and compare it to
// other levels
const InfoLevel int = 70

// DebugLevelName is the name of the DEBUG log level
const DebugLevelName string = "DEBUG"

// DebugLevel is used to identify the DEBUG log level and compare it to
// other levels
const DebugLevel int = 60

// Levels maps level priorities to level names
var Levels map[int]string = map[int]string{
	FatalLevel: FatalLevelName,
	ErrorLevel: ErrorLevelName,
	WarnLevel:  WarnLevelName,
	InfoLevel:  InfoLevelName,
	DebugLevel: DebugLevelName,
}
