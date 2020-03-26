package glog

// support:
// 1. redirect output
// 2. colors
// 3. prefix
// 4. log level

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level uint8

const (
	EMERG   = iota // 0  usually user did't need this level.
	ALERT          // 1
	CRIT           // 2
	ERR            // 3
	WARNING        // 4
	NOTICE         // 5
	INFO           // 6
	DEBUG          // 7
)

func (l Level) String() string {
	switch l {
	case EMERG:
		return "ERERG"
	case ALERT:
		return "ALERT"
	case CRIT:
		return "CRIT"
	case ERR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "INFO"
	}
}

func parseLevel(lstr string) Level {
	switch lstr {
	case "emergency", "emerg":
		return EMERG
	case "alert":
		return ALERT
	case "critical", "crit", "criti":
		return CRIT
	case "error", "err":
		return ERR
	case "warning", "warn":
		return WARNING
	case "notice":
		return NOTICE
	case "info":
		return INFO
	case "debug":
		return DEBUG
	}
	return INFO
}

type Logger struct {
	Out    io.Writer
	prefix string
	logger *log.Logger
}

func New(prefixs ...string) *Logger {
	prefix := ""
	if len(prefixs) > 0 {
		prefix = prefixs[0]
	}
	return &Logger{
		Out:    os.Stdout,
		prefix: prefix,
		logger: &log.Logger{},
	}
}

func (l *Logger) SetOutput(output *os.File) *Logger {
	l.Out = output
	return l
}

// shutdown logger.
func (l *Logger) Enable(level Level) {

}

func (l *Logger) Emerg(msgs ...interface{}) {}
func (l *Logger) Alert(msgs ...interface{}) {}
func (l *Logger) Crit(msgs ...interface{})  {}
func (l *Logger) Err(msgs ...interface{}) {
	l.output(ERR, msgs...)
}
func (l *Logger) Warning(msgr ...interface{}) {}
func (l *Logger) Notice(msgs ...interface{})  {}
func (l *Logger) Info(msgs ...interface{}) {
	l.output(INFO, msgs...)
}
func (l *Logger) Debug(msgs ...interface{}) {}

func (l *Logger) Time(prefix string)    {}
func (l *Logger) TimeEnd(prefix string) {}

func (l *Logger) output(level Level, msgs ...interface{}) {

	mark := "[" + level.String() + "]"
	infos := make([]interface{}, 0, len(msgs)+2)

	infos = append(infos, mark)
	infos = append(infos, l.prefix)
	infos = append(infos, msgs...)

	if level > ERR {
		fmt.Fprintln(l.Out, infos...)
	} else {
		fmt.Fprintln(os.Stderr, infos...)
	}
}
