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
	LOGLEVEL_ERROR
	LOGLEVEL_WARN
	LOGLEVEL_CONFIG
	LOGLEVEL_INFO
	LOGLEVEL_DEBUG
	nr_loglevels // levels count
)

// flags
const (
	FLAG_Ldate         = log.Ldate         // the date: 2009/01/23
	FLAG_Ltime         = log.Ltime         // the time: 01:23:23
	FLAG_Lmicroseconds = log.Lmicroseconds // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	FLAG_Llongfile     = log.Llongfile     // full file name and line number: /a/b/c/d.go:23
	FLAG_Lshortfile    = log.Lshortfile    // final file name element and line number: d.go:23. overrides Llongfile
	FLAG_LstdFlags     = log.LstdFlags     // initial values for the standard logger
)

// used to print file and line information 
const calldepth = 2

type Logger struct {
	logLevel int
	logger   *log.Logger
}

type NullWriter struct{}

func (f *NullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

var logLevelToName map[int]string

func init() {
	logLevelToName = make(map[int]string, nr_loglevels)
	logLevelToName[LOGLEVEL_DEBUG] = "[Debug] "
	logLevelToName[LOGLEVEL_INFO] = "[Info] "
	logLevelToName[LOGLEVEL_CONFIG] = "[Config] "
	logLevelToName[LOGLEVEL_WARN] = "[Warning] "
	logLevelToName[LOGLEVEL_ERROR] = "[Error] "
	logLevelToName[LOGLEVEL_FATAL] = "[Fatal] "
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.logLevel < LOGLEVEL_DEBUG {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_DEBUG]+fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.logLevel < LOGLEVEL_DEBUG {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_DEBUG]+fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Info(v ...interface{}) {
	if l.logLevel < LOGLEVEL_INFO {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_INFO]+fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.logLevel < LOGLEVEL_INFO {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_INFO]+fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Config(v ...interface{}) {
	if l.logLevel < LOGLEVEL_CONFIG {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_CONFIG]+fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Configf(format string, v ...interface{}) {
	if l.logLevel < LOGLEVEL_CONFIG {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_CONFIG]+fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Warn(v ...interface{}) {
	if l.logLevel < LOGLEVEL_WARN {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_WARN]+fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.logLevel < LOGLEVEL_WARN {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_WARN]+fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Error(v ...interface{}) {
	if l.logLevel < LOGLEVEL_ERROR {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_ERROR]+fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.logLevel < LOGLEVEL_ERROR {
		return
	}
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_ERROR]+fmt.Sprintf(format, v...))
	checkErr(err)
}

func (l *Logger) Fatal(v ...interface{}) {
	defer os.Exit(1)
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_FATAL]+fmt.Sprint(v...))
	checkErr(err)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	defer os.Exit(1)
	err := l.logger.Output(calldepth, logLevelToName[LOGLEVEL_FATAL]+fmt.Sprintf(format, v...))
	checkErr(err)
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
	if writer == nil {
		writer = &NullWriter{}
	}

	if logLevel < LOGLEVEL_FATAL {
		logLevel = LOGLEVEL_FATAL
	} else if logLevel > LOGLEVEL_DEBUG {
		logLevel = LOGLEVEL_DEBUG
	}
	ret.logLevel = logLevel

	ret.logger = log.New(writer, prefix+" ", flag)
	return ret
}
