package logging

import (
	"context"
	"log/slog"
)

var RequestIdHeader string = "X-Request-Id"

type LoggingHandler struct {
	Environment string
	slog.Handler
}

// Retrieves a string from the provided context with the given key
func retrieveCtxString(ctx context.Context, key any) string {
	if val, ok := ctx.Value(key).(string); ok {
		return val
	}

	return "unknown"
}

func New(opts ...func(*LoggingHandler)) *LoggingHandler {
	handler := &LoggingHandler{}

	for _, fn := range opts {
		fn(handler)
	}

	return handler
}

func WithEnvironment(env string) func(*LoggingHandler) {
	return func(handler *LoggingHandler) {
		handler.Environment = env
	}
}

// Override the base Handle function to insert our attrs
func (handler *LoggingHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(slog.String("request-id", retrieveCtxString(ctx, RequestIdHeader)))
	r.AddAttrs(slog.String("environment", retrieveCtxString(ctx, "environment")))

	return handler.Handler.Handle(ctx, r)
}
