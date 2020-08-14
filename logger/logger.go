package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"io"
	"path/filepath"
	"strconv"
)

// https://mp.weixin.qq.com/s/hZzcce4S4D0B6tXeorb8-A
// https://github.com/polaris1119/go-echo-example/blob/master/pkg/logger/logger.go

var logDefaultHeader = map[string]string{
	"time":   "${time_rfc3339_name}",
	"level":  "${level}",
	"prefix": "${prefix}",
	"file":   "${file}",
	"line":   "${line}",
}

func init() {
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
}

type Logger struct {
	*log.Logger
	ZeroLog zerolog.Logger
}

var _ echo.Logger = &Logger{}

func New(writer io.Writer) *Logger {
	l := &Logger{
		Logger:  log.New("-"),
		ZeroLog: zerolog.New(writer).With().Caller().Timestamp().Logger(),
	}
	// default level is ERROR, change it
	l.SetLevel(log.INFO)

	l.Logger.SetOutput(writer)

	return l
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.Logger.SetOutput(writer)
	l.ZeroLog.Output(writer)
}

func (l *Logger) SetLevel(level log.Lvl) {
	l.Logger.SetLevel(level)
	if level == log.OFF {
		l.ZeroLog = l.ZeroLog.Level(zerolog.Disabled)
	} else {
		zeroLeve := int8(level) - 1
		l.ZeroLog = l.ZeroLog.Level(zerolog.Level(zeroLeve))
	}
}
