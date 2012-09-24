// golog provides a convenient API for logging with golang.
// You can use log level in your program with golog.
package golog

import (
	"io"
	"log"
	"fmt"
	"os"
)

// log levels
const (
	LOGLEVEL_FATAL = iota
	LOGLEVEL_ALERT
	LOGLEVEL_ERROR
	LOGLEVEL_WARN
	LOGLEVEL_CONFIG
	LOGLEVEL_INFO
	LOGLEVEL_DEBUG
	NR_LOGLEVELS
)

// flags
const (
	FLAG_Ldate = log.Ldate     // the date: 2009/01/23
	FLAG_Ltime = log.Ltime     // the time: 01:23:23
	FLAG_Lmicroseconds = log.Lmicroseconds  // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	FLAG_Llongfile = log.Llongfile       // full file name and line number: /a/b/c/d.go:23
	FLAG_Lshortfile = log.Lshortfile        // final file name element and line number: d.go:23. overrides Llongfile
	FLAG_LstdFlags = log.LstdFlags // initial values for the standard logger
)

// used to print file and line information 
const calldepth = 2

type Logger struct {
	logLevel int
	loggers  []*log.Logger
	prefix   string
	writer   io.Writer
}

type NullWriter struct{}

func (f *NullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (l *Logger) Debug(v ...interface{}) {
	err := l.loggers[LOGLEVEL_DEBUG].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_DEBUG].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Info(v ...interface{}) {
	err := l.loggers[LOGLEVEL_INFO].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_INFO].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Config(v ...interface{}) {
	err := l.loggers[LOGLEVEL_CONFIG].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Configf(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_CONFIG].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Warn(v ...interface{}) {
	err := l.loggers[LOGLEVEL_WARN].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_WARN].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Error(v ...interface{}) {
	err := l.loggers[LOGLEVEL_ERROR].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_ERROR].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Alert(v ...interface{}) {
	err := l.loggers[LOGLEVEL_ALERT].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Alertf(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_ALERT].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Fatal(v ...interface{}) {
	err := l.loggers[LOGLEVEL_FATAL].Output(calldepth, fmt.Sprint(v...))
	checkErr(err)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	err := l.loggers[LOGLEVEL_FATAL].Output(calldepth, fmt.Sprintf(format, v...))
	checkErr(err)
	os.Exit(1)
}

var logLevelToName map[int]string

func init() {
	logLevelToName = make(map[int]string, NR_LOGLEVELS)
	logLevelToName[LOGLEVEL_DEBUG] = "[Debug]"
	logLevelToName[LOGLEVEL_INFO] = "[Info]"
	logLevelToName[LOGLEVEL_CONFIG] = "[Config]"
	logLevelToName[LOGLEVEL_WARN] = "[Warning]"
	logLevelToName[LOGLEVEL_ERROR] = "[Error]"
	logLevelToName[LOGLEVEL_ALERT] = "[Alert]"
	logLevelToName[LOGLEVEL_FATAL] = "[Fatal]"
}

// if @writer is nil, log nothing
// if @writer is file, you should guarantee that you have write permition
// @prefix will appear in the header of every line
// @logLevel should be one of golog.LOGLEVEL_*, for example golog.LOGLEVEL_DEBUG
// @flag decide to append extra information, such as date, time, file name, line and so on.
// if you want to output date, flag should be golog.FLAG_Ldate, to output time, flag should be golog.FLAG_Ltime.
// to output date and time, then golog.FLAG_Ldate | golog.FLAG_Ltime
func NewLogger(writer io.Writer, prefix string, logLevel int, flag int) *Logger {
	ret := new(Logger)
	ret.loggers = make([]*log.Logger, NR_LOGLEVELS)
	if writer == nil {
		ret.writer = &NullWriter{}
	} else {
		ret.writer = writer
	}
	ret.prefix = prefix
	ret.setLogLevel(logLevel, flag)
	return ret
}

func (l *Logger) setLogLevel(logLevel int, flag int) {
	if logLevel > LOGLEVEL_DEBUG {
		logLevel = LOGLEVEL_DEBUG
	} else if logLevel < LOGLEVEL_FATAL {
		logLevel = LOGLEVEL_FATAL - 1
	}
	l.logLevel = logLevel
	for i := LOGLEVEL_FATAL; i <= logLevel; i++ {
		l.loggers[i] = log.New(l.writer, l.prefix+logLevelToName[i]+" ", flag)
	}
	nullwriter := &NullWriter{}
	for i := logLevel + 1; i < NR_LOGLEVELS; i++ {
		l.loggers[i] = log.New(nullwriter, l.prefix+logLevelToName[i]+" ", flag)
	}
}