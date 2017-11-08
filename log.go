package log

import (
	"os"

	"github.com/mdeheij/logwrap/go-logging"
)

var stdoutFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05}: %{message} %{color:reset}`,
)

//log is an exported logging variable
var log = logging.MustGetLogger("example")

func init() {
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	backendStdOutFormatter := logging.NewBackendFormatter(stdout, stdoutFormat)

	logging.SetBackend(backendStdOutFormatter)
}

//Debug See log.Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

//Debugf See log.Debugf
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

//Info See log.Info
func Info(args ...interface{}) {
	log.Info(args...)
}

//Infof See log.Infof
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

//Notice See log.Notice
func Notice(args ...interface{}) {
	log.Notice(args...)
}

//Noticef See log.Noticef
func Noticef(format string, args ...interface{}) {
	log.Noticef(format, args...)
}

//Warning See log.Warning
func Warning(args ...interface{}) {
	log.Warning(args...)
}

//Warningf See log.Warningf
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

//Error See log.Error
func Error(args ...interface{}) {
	log.Error(args...)
}

//Errorf See log.Errorf
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

//Critical See log.Critical
func Critical(args ...interface{}) {
	log.Critical(args...)
}

//Criticalf See log.Criticalf
func Criticalf(format string, args ...interface{}) {
	log.Criticalf(format, args...)
}

//Panic See log.Panic
func Panic(args ...interface{}) {
	log.Panic(args...)
}

//Panicf See log.Panicf
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

//Fatal See log.Fatal
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

//Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}