package golog_test

import (
	"github.com/Noah-Huppert/golog"
)

// Shows how to use Go Log with no customization.
func Example_basic() {
	// logger will print normal messages to stdout and errors to stderr
	logger := golog.NewStdLogger("basic-example")

	// Log messages for each log level
	logger.Debug("hello debug")
	logger.Debugf("hello %s", "debug")

	logger.Info("hello info")
	logger.Infof("hello %s", "info")

	logger.Warn("hello warn")
	logger.Warnf("hello %s", "warn")

	logger.Error("hello error")
	logger.Errorf("hello %s", "error")

	logger.Fatal("hello fatal")
	logger.Fatalf("hello %s", "fatal")

	// Output:
	// basic-example [DEBUG] hello debug
	// basic-example [DEBUG] hello debug
	// basic-example [INFO] hello info
	// basic-example [INFO] hello info
	// basic-example [WARN] hello warn
	// basic-example [WARN] hello warn
	// basic-example [ERROR] hello error
	// basic-example [ERROR] hello error
	// basic-example [FATAL] hello fatal
	// basic-example [FATAL] hello fatal
}

// Shows how to only display messages from certain log levels
func Example_levels() {
	// Configure logger to only display error messages or greater
	logger := golog.NewStdLogger("levels-example")
	logger.SetLevel(golog.ErrorLevel)

	// Log messages which will not be shown because they are below the
	// specified log level
	logger.Debug("I will not be shown b/c I am a debug message")
	logger.Info("I am just an info message show I will not be shown")
	logger.Warn("Due to the fact that I am a warn message I will not be displayed")

	// Log message that will show
	logger.Error("Error log messages will show")
	logger.Fatal("I am a fatal message so I will be displayed")

	// Output:
	// levels-example [ERROR] Error log messages will show
	// levels-example [FATAL] I am a fatal message so I will be displayed
}

// Shows how to customize the log output format
func Example_format() {
	// Configure logger with special format
	logger := golog.NewStdLogger("basic-example")
	logger.SetFormatTmpl("name={{ .Name }} level={{ .Level }} msg={{ .Msg }}")

	// Log messages for each log level
	logger.Debug("hello debug")
	logger.Debugf("hello %s", "debug")

	logger.Info("hello info")
	logger.Infof("hello %s", "info")

	logger.Warn("hello warn")
	logger.Warnf("hello %s", "warn")

	logger.Error("hello error")
	logger.Errorf("hello %s", "error")

	logger.Fatal("hello fatal")
	logger.Fatalf("hello %s", "fatal")

	// Output:
	// name=basic-example level=DEBUG msg=hello debug
	// name=basic-example level=DEBUG msg=hello debug
	// name=basic-example level=INFO msg=hello info
	// name=basic-example level=INFO msg=hello info
	// name=basic-example level=WARN msg=hello warn
	// name=basic-example level=WARN msg=hello warn
	// name=basic-example level=ERROR msg=hello error
	// name=basic-example level=ERROR msg=hello error
	// name=basic-example level=FATAL msg=hello fatal
	// name=basic-example level=FATAL msg=hello fatal
}
