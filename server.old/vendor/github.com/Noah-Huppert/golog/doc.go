// Go Log is a simple to use logging library inspired by the standard `log`
// package.
//
// It provides the: Debug, Info, Warn, Error, and Fatal log levels.
//
// Go Log is easy to setup:
//
//   // logger will print normal messages to stdout and errors to stderr
//   logger := golog.NewStdLogger("example")
//
// The logging API should be familiar to those who have used the standard
// `fmt` and `log` packages.
//
//   logger.Debug("hello debug")
//   logger.Debugf("hello %s", "debug")
//
//   logger.Info("hello info")
//   logger.Infof("hello %s", "info")
//
//   logger.Warn("hello warn")
//   logger.Warnf("hello %s", "warn")
//
//   logger.Error("hello error")
//   logger.Errorf("hello %s", "error")
//
//   logger.Fatal("hello fatal")
//   logger.Fatalf("hello %s", "fatal")
//
// You can configure Go Log to only show messages of certain importance
//
//   logger.SetLevel(golog.DebugLevel)
//
//   logger.SetLevel(golog.InfoLevel)
//
//   logger.SetLevel(golog.WarnLevel)
//
//   logger.SetLevel(golog.ErrorLevel)
//
//   logger.SetLevel(golog.FatalLevel)
//
// Log output format can be configured with Go templates
//
//  logger.SetFormatTmpl("name={{ .Name }} level={{ .Level }} msg={{ .Msg }}")
//
package golog
