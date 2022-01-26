package logging

import (
	"context"
)

type logKey string

const contextLog logKey = "contextLog"

func WithValue(ctx context.Context, logger *ZapLogger) context.Context {
	return context.WithValue(ctx, contextLog, logger)
}

func GetContextLogger(ctx context.Context) *ZapLogger {
	return ctx.Value(contextLog).(*ZapLogger)
}
