package logger

import (
	"fmt"
	"log/slog"
	"os"
	"sync/atomic"
)

var defaultLogger atomic.Pointer[Logger]

func init() {
	defaultLogger.Store(NewLogger(slog.Default().WithGroup("my-app")))
}

// Default returns the default [Logger].
func Default() *Logger { return defaultLogger.Load() }

type Logger struct {
	log *slog.Logger
}

func NewLogger(logger *slog.Logger) *Logger {
	return &Logger{log: logger}
}

// Fatal is equivalent to [Print] followed by a call to [os.Exit](1).
func (l *Logger) Fatal(v ...any) {
	l.log.Error(fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to [Printf] followed by a call to [os.Exit](1).
func (l *Logger) Fatalf(format string, v ...any) {
	l.log.Error(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Error(v ...any) {
	l.log.Error(fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...any) {
	l.log.Error(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...any) {
	l.log.Info(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...any) {
	l.log.Info(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...any) {
	l.log.Warn(fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...any) {
	l.log.Warn(fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...any) {
	l.log.Debug(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...any) {
	l.log.Debug(fmt.Sprintf(format, v...))
}

func Fatal(v ...any) {
	Default().Fatal(v)
}

// Fatalf is equivalent to [Printf] followed by a call to [os.Exit](1).
func Fatalf(format string, v ...any) {
	Default().Fatalf(format, v)
}

func Error(v ...any) {
	Default().Error(v)
}

func Errorf(format string, v ...any) {
	Default().Errorf(format, v)
}

func Info(v ...any) {
	Default().Info(v)
}

func Infof(format string, v ...any) {
	Default().Infof(format, v)
}

func Warn(v ...any) {
	Default().Warn(v)
}

func Warnf(format string, v ...any) {
	Default().Warnf(format, v)
}

func Debug(v ...any) {
	Default().Debug(v)
}

func Debugf(format string, v ...any) {
	Default().Debugf(format, v)
}
