package main

import (
	"fmt"
	"strings"
	"time"
)

type LogLevel uint8

const (
	logLevelDebug LogLevel = iota
	logLevelInfo
	logLevelWarn
	logLevelError
	logLevelFatal
)

type Logger struct {
	logLevel   LogLevel
	logColors  map[LogLevel]string
	timeFormat string
	timeColor  string
	resetColor string
}

var logger Logger = Logger{
	logLevel: logLevelDebug,
	logColors: map[LogLevel]string{
		logLevelDebug: "\033[36m", // Cyan
		logLevelInfo:  "\033[32m", // Green
		logLevelWarn:  "\033[33m", // Yellow
		logLevelError: "\033[31m", // Red
		logLevelFatal: "\033[35m", // Magenta
	},
	timeFormat: "15:04:05",
	timeColor:  "\033[90m", // Bright Black / Gray
	resetColor: "\033[0m",
}

func (L Logger) Debug(a ...any) {
	if L.logLevel > logLevelDebug {
		return
	}

	args := append([]any{
		L.timeColor + time.Now().Format(L.timeFormat) + L.resetColor,
		L.logColors[logLevelDebug] + "DBG" + L.resetColor,
	}, a...)
	fmt.Println(args...)
}

func (L Logger) Debugf(format string, a ...any) {
	if L.logLevel > logLevelDebug {
		return
	}

	fmt.Printf(
		L.timeColor+time.Now().Format(L.timeFormat)+L.resetColor+
			" "+L.logColors[logLevelDebug]+"DBG"+L.resetColor+" "+format,
		a...,
	)
}

func (L Logger) Info(a ...any) {
	if L.logLevel > logLevelInfo {
		return
	}

	args := append([]any{
		L.timeColor + time.Now().Format(L.timeFormat) + L.resetColor,
		L.logColors[logLevelInfo] + "INF" + L.resetColor,
	}, a...)
	fmt.Println(args...)
}

func (L Logger) Infof(format string, a ...any) {
	if L.logLevel > logLevelInfo {
		return
	}

	fmt.Printf(
		L.timeColor+time.Now().Format(L.timeFormat)+L.resetColor+
			" "+L.logColors[logLevelInfo]+"INF"+L.resetColor+" "+format,
		a...,
	)
}

func (L Logger) Warn(a ...any) {
	if L.logLevel > logLevelWarn {
		return
	}

	args := append([]any{
		L.timeColor + time.Now().Format(L.timeFormat) + L.resetColor,
		L.logColors[logLevelWarn] + "WRN" + L.resetColor,
	}, a...)
	fmt.Println(args...)
}

func (L Logger) Warnf(format string, a ...any) {
	if L.logLevel > logLevelWarn {
		return
	}

	fmt.Printf(
		L.timeColor+time.Now().Format(L.timeFormat)+L.resetColor+
			" "+L.logColors[logLevelWarn]+"WRN"+L.resetColor+" "+format,
		a...,
	)
}

func (L Logger) Error(a ...any) {
	if L.logLevel > logLevelError {
		return
	}

	args := append([]any{
		L.timeColor + time.Now().Format(L.timeFormat) + L.resetColor,
		L.logColors[logLevelError] + "ERR" + L.resetColor,
	}, a...)
	fmt.Println(args...)
}

func (L Logger) Errorf(format string, a ...any) {
	if L.logLevel > logLevelError {
		return
	}

	fmt.Printf(
		L.timeColor+time.Now().Format(L.timeFormat)+L.resetColor+
			" "+L.logColors[logLevelError]+"ERR"+L.resetColor+" "+format,
		a...,
	)
}

func (L Logger) Fatal(a ...any) {
	if L.logLevel > logLevelFatal {
		return
	}

	args := append([]any{
		L.timeColor + time.Now().Format(L.timeFormat) + L.resetColor,
		L.logColors[logLevelFatal] + "FTL" + L.resetColor,
	}, a...)
	fmt.Println(args...)
}

func (L Logger) Fatalf(format string, a ...any) {
	if L.logLevel > logLevelFatal {
		return
	}

	fmt.Printf(
		L.timeColor+time.Now().Format(L.timeFormat)+L.resetColor+
			" "+L.logColors[logLevelFatal]+"FTL"+L.resetColor+" "+format,
		a...,
	)
}

func logLevelFromString(level string) LogLevel {
	switch strings.ToLower(level) {
	case "debug", "dgb", "d":
		return logLevelDebug
	case "info", "inf", "i":
		return logLevelInfo
	case "warn", "wrn", "w":
		return logLevelWarn
	case "error", "err", "e":
		return logLevelError
	case "fatal", "ftl", "f":
		return logLevelFatal
	default:
		return logLevelDebug
	}
}
