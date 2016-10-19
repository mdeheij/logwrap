package log

import (
	"os"

	"github.com/op/go-logging"
)

var stdoutFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05}: %{message}`,
)

//log is an exported logging variable
var log = logging.MustGetLogger("example")

func init() {
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	backendStdOutFormatter := logging.NewBackendFormatter(stdout, stdoutFormat)

	logging.SetBackend(backendStdOutFormatter)
}

//Debugf  See log.Debugf
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

//Debug  See log.Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

//Info  See log.Info
func Info(args ...interface{}) {
	log.Info(args...)
}

//Println  See log.Info
func Println(args ...interface{}) {
	log.Info(args...)
}

//Notice  See log.Notice
func Notice(args ...interface{}) {
	log.Notice(args...)
}

//Warning  See log.Warning
func Warning(args ...interface{}) {
	log.Warning(args...)
}

//Error  See log.Error
func Error(args ...interface{}) {
	log.Error(args...)
}

//Critical  See log.Critical
func Critical(args ...interface{}) {
	log.Critical(args...)
}

//Panic  See log.Panic
func Panic(args ...interface{}) {
	log.Panic(args...)
}

//Fatal  See log.Fatal
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
