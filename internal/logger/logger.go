package logger

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jamillosantos/logctx"
	"go.uber.org/zap"
)

func GetLogger(ctx context.Context) *zap.Logger {
	return logctx.From(ctx)
}

func GetLoggerComponent(ctx context.Context, name string) *zap.Logger {
	return GetLogger(ctx).With(zap.String("component", name))
}

func CreateSpan(ctx context.Context, componentName *string, spanName string) context.Context {
	uuid, _ := uuid.NewV7()
	if componentName != nil {
		spanName = fmt.Sprintf("%s::%s", *componentName, spanName)
	}
	return logctx.WithFields(ctx, zap.String("span_id", uuid.String()), zap.String("span", spanName))
}

func CreateTrace(ctx context.Context) context.Context {
	uuid, _ := uuid.NewV7()
	return logctx.WithFields(ctx, zap.String("trace_id", uuid.String()))
}
