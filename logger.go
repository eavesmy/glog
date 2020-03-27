package glog

// support:
// 1. redirect output
// 2. colors
// 3. prefix
// 4. log level

import (
	// "fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
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
	prefix string
	level  Level
	logger *log.Logger
	unable map[Level]bool
	sinces map[string]time.Time
}

func New(prefixs ...string) *Logger {
	prefix := ""
	if len(prefixs) > 0 {
		prefix = prefixs[0]
	}

	return &Logger{
		level:  INFO,
		prefix: prefix,
		logger: log.New(os.Stdout, "", log.Lshortfile),
		unable: make(map[Level]bool),
		sinces: make(map[string]time.Time),
	}
}

func (l *Logger) SetOutput(output io.Writer) *Logger {
	l.logger.SetOutput(output)
	return l
}

func (l *Logger) SetLevel(level Level) *Logger {
	l.level = level
	return l
}

// shutdown logger.
func (l *Logger) Unable(c string) {
	for _, str_level := range strings.Split(c, ",") {
		level := parseLevel(str_level)
		l.unable[level] = true
	}
}

func (l *Logger) Emerg(msgs ...interface{}) {
	l.output(EMERG, msgs...)
}
func (l *Logger) Alert(msgs ...interface{}) {
	l.output(ALERT, msgs...)
}
func (l *Logger) Crit(msgs ...interface{}) {
	l.output(CRIT, msgs...)
}
func (l *Logger) Err(msgs ...interface{}) {
	l.output(ERR, msgs...)
}
func (l *Logger) Warning(msgs ...interface{}) {
	l.output(WARNING, msgs...)
}
func (l *Logger) Notice(msgs ...interface{}) {
	l.output(NOTICE, msgs...)
}
func (l *Logger) Info(msgs ...interface{}) {
	l.output(INFO, msgs...)
}

func (l *Logger) Debug(msgs ...interface{}) {
	l.output(DEBUG, msgs...)
}
func (l *Logger) Time(prefix string) {
	l.sinces[prefix] = time.Now()
}

func (l *Logger) Println(v ...interface{}) {
	l.output(l.level, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.SetPrefix("[" + l.level.String() + "] ")
	l.logger.Printf(format, v...)
}

// Do not use this method.
// Is a bad idea to write this.
func (l *Logger) Print(v ...interface{}) {
	l.logger.SetPrefix("[" + l.level.String() + "] ")
	l.logger.Print(v...)
}

func (l *Logger) TimeEnd(prefix string) {
	if t, exists := l.sinces[prefix]; exists {
		l.Debug(prefix, time.Since(t))
	}
	l.Err("No this prefix", prefix)
}

func (l *Logger) output(level Level, msgs ...interface{}) {

	if l.unable[level] {
		return
	}

	mark := "[" + level.String() + "] "
	infos := make([]interface{}, 0, len(msgs)+2)

	infos = append(infos, l.prefix)
	infos = append(infos, msgs...)

	l.logger.SetPrefix(mark)

	l.logger.Println(infos...)
}
