// Package lgl provides a very simple logger with log level support
//
// Supported log levels: ERROR, DEBUG, WARNING, INFO, TRACE.
// Log level is default to ERROR if Set/SetAll wasn't called in main().
// All the output functions make sure a newline is at the end of the output.
package lgl

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	LError LogLevel = iota
	LWarning
	LInfo
	LDebug
	LTrace
)
const flag = log.LstdFlags | log.Lshortfile

var (
	traceLogger   *log.Logger
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   = log.New(os.Stderr, "E: ", flag)
)

// Set sets log level and output writers
//
// Call this function in top function (normally it is main),
// If not, only error message is written to stderr.
func Set(l LogLevel, out, err io.Writer) {
	switch l {
	case LTrace:
		traceLogger = log.New(out, "T: ", flag)
		fallthrough
	case LDebug:
		debugLogger = log.New(out, "D: ", flag)
		fallthrough
	case LInfo:
		infoLogger = log.New(out, "I: ", flag)
		fallthrough
	case LWarning:
		warningLogger = log.New(err, "W: ", flag)
		fallthrough
	default:
		errorLogger = log.New(err, "E: ", flag)
	}
}

// SetLogLeveL sets log level
//
// Call this function in top function (normally it is main),
// If not, only error message is output.
func SetLogLeveL(l LogLevel) {
	Set(l, os.Stdout, os.Stderr)
}

// Trace works as log.Println()
func Trace(v ...interface{}) {
	if traceLogger == nil {
		return
	}
	traceLogger.Output(2, fmt.Sprint(v...))
}

// Trace works as log.Printf()
func Tracef(format string, v ...interface{}) {
	if traceLogger == nil {
		return
	}
	traceLogger.Output(2, fmt.Sprintf(format, v...))
}

// Debug works as log.Println()
func Debug(v ...interface{}) {
	if debugLogger == nil {
		return
	}
	debugLogger.Output(2, fmt.Sprint(v...))
}

// Debugf works as log.Printf()
func Debugf(format string, v ...interface{}) {
	if debugLogger == nil {
		return
	}
	debugLogger.Output(2, fmt.Sprintf(format, v...))
}

// Info works as log.Println()
func Info(v ...interface{}) {
	if infoLogger == nil {
		return
	}
	infoLogger.Output(2, fmt.Sprint(v...))
}

// Infof works as log.Printf()
func Infof(format string, v ...interface{}) {
	if infoLogger == nil {
		return
	}
	infoLogger.Output(2, fmt.Sprintf(format, v...))
}

// Warning works as log.Println()
func Warning(v ...interface{}) {
	if warningLogger == nil {
		return
	}
	warningLogger.Output(2, fmt.Sprint(v...))
}

// Warningf works as log.Printf()
func Warningf(format string, v ...interface{}) {
	if warningLogger == nil {
		return
	}
	warningLogger.Output(2, fmt.Sprintf(format, v...))
}

// Error works as log.Println()
func Error(v ...interface{}) {
	if errorLogger == nil {
		return
	}
	errorLogger.Output(2, fmt.Sprint(v...))
}

// Errorf works as log.Printf()
func Errorf(format string, v ...interface{}) {
	if errorLogger == nil {
		return
	}
	errorLogger.Output(2, fmt.Sprintf(format, v...))
}
