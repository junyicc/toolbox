package logging

import (
	"fmt"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
	level  zapcore.Level
	writer io.Writer
}

func (l *ZapLogger) SetLogger(logger *zap.Logger) {
	l.logger = logger
}

func (l *ZapLogger) GetLogger() *zap.Logger {
	return l.logger
}

func (l *ZapLogger) Debug(msg string, args ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Debug(msg, args...)
}

func (l *ZapLogger) Info(msg string, args ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Info(msg, args...)
}

func (l *ZapLogger) Warn(msg string, args ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Warn(msg, args...)
}

func (l *ZapLogger) Error(msg string, args ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Error(msg, args...)
}

func (l *ZapLogger) Panic(msg string, args ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Panic(msg, args...)
}
func (l *ZapLogger) Fatal(msg string, args ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Fatal(msg, args...)
}

// Printf implements gorm.Logger.Writer interface
func (l *ZapLogger) Printf(msg string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Debug(fmt.Sprintf(msg, args...))
}

func (l *ZapLogger) Writer() io.Writer {
	return l.writer
}

func (l *ZapLogger) Sync() {
	if l.logger == nil {
		return
	}
	if err := l.logger.Sync(); err != nil {
		panic(err)
	}
}

func (l *ZapLogger) Level() zapcore.Level {
	return l.level
}

func (l *ZapLogger) With(fields ...zap.Field) *ZapLogger {
	if len(fields) == 0 {
		return l
	}
	zl := *l
	zl.logger = l.logger.With(fields...)
	return &zl
}

// DefaultLogger is a global default ZapLogger
var DefaultLogger = &ZapLogger{
	logger: zap.NewNop(),
	level:  zap.FatalLevel + 1,
	writer: io.Discard,
}

func Debug(msg string, args ...zap.Field) {
	DefaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...zap.Field) {
	DefaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...zap.Field) {
	DefaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...zap.Field) {
	DefaultLogger.Error(msg, args...)
}

func Panic(msg string, args ...zap.Field) {
	DefaultLogger.Panic(msg, args...)
}

func Fatal(msg string, args ...zap.Field) {
	DefaultLogger.Fatal(msg, args...)
}

func Writer() io.Writer {
	return DefaultLogger.Writer()
}

func Sync() {
	DefaultLogger.Sync()
}

func Level() zapcore.Level {
	return DefaultLogger.Level()
}

// With creates a child logger and adds structured context to it
func With(fields ...zap.Field) *ZapLogger {
	return DefaultLogger.With(fields...)
}
