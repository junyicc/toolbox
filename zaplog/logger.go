package zaplog

import (
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, v ...zap.Field)
	Info(msg string, v ...zap.Field)
	Warn(msg string, v ...zap.Field)
	Error(msg string, v ...zap.Field)
	Panic(msg string, v ...zap.Field)
	Fatal(msg string, v ...zap.Field)
	Writer() io.Writer
	Flush()
}

type ZapLogger struct {
	logger *zap.Logger
	writer io.Writer
}

func (l *ZapLogger) Debug(msg string, v ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Debug(msg, v...)
}

func (l *ZapLogger) Info(msg string, v ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Info(msg, v...)
}

func (l *ZapLogger) Warn(msg string, v ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Warn(msg, v...)
}

func (l *ZapLogger) Error(msg string, v ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Error(msg, v...)
}

func (l *ZapLogger) Panic(msg string, v ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Panic(msg, v...)
}
func (l *ZapLogger) Fatal(msg string, v ...zap.Field) {
	if l.logger == nil {
		return
	}
	l.logger.Fatal(msg, v...)
}

func (l *ZapLogger) Writer() io.Writer {
	return l.writer
}

func (l *ZapLogger) Flush() {
	if err := l.logger.Sync(); err != nil {
		panic(err)
	}
}

// zaplogger is a global default ZapLogger
// implement Logger interface
var zaplogger = &ZapLogger{
	logger: nil,
	writer: os.Stdout,
}

var once sync.Once

// init default ZapLogger
func init() {
	once.Do(func() {
		initLogger()
	})
}

func initLogger() {
	// print all level log
	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	// writers
	consoleWriter := zapcore.Lock(os.Stdout)
	// encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleWriter, allPriority),
	)

	zaplogger = &ZapLogger{
		logger: zap.New(core, zap.AddCaller()),
		writer: consoleWriter,
	}
}

func Debug(msg string, v ...zap.Field) {
	zaplogger.Debug(msg, v...)
}

func Info(msg string, v ...zap.Field) {
	zaplogger.Info(msg, v...)
}

func Warn(msg string, v ...zap.Field) {
	zaplogger.Warn(msg, v...)
}

func Error(msg string, v ...zap.Field) {
	zaplogger.Error(msg, v...)
}

func Panic(msg string, v ...zap.Field) {
	zaplogger.Panic(msg, v...)
}

func Fatal(msg string, v ...zap.Field) {
	zaplogger.Fatal(msg, v...)
}

func Writer() io.Writer {
	return zaplogger.Writer()
}

func Flush() {
	zaplogger.Flush()
}

// With creates a child logger and adds structured context to it
func With(v ...zap.Field) *ZapLogger {
	return &ZapLogger{
		logger: zaplogger.logger.With(v...),
		writer: zaplogger.writer,
	}
}

// WithWriter clones the ZapLogger with a new writer
func WithWriter(w io.Writer) *ZapLogger {
	return &ZapLogger{
		logger: zaplogger.logger.WithOptions(zap.AddCaller()),
		writer: w,
	}
}
