package logging

import (
	"context"
	"log/slog"
	"os"
)

var RequestIdHeader string = "X-Request-Id"

type LoggingHandler struct {
	Environment string
	slog.Handler
}

func NewHandler(opts ...func(*LoggingHandler)) *LoggingHandler {
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

func WithBaseHandler(structured bool, level slog.Level, addSource bool) func(*LoggingHandler) {
	if structured {
		return WithBaseJsonHandler(level, addSource)
	} else {
		return WithBaseTextHandler(level, addSource)
	}
}

func WithBaseJsonHandler(level slog.Level, addSource bool) func(*LoggingHandler) {
	return func(handler *LoggingHandler) {
		handler.Handler = slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level:     level,
				AddSource: addSource,
			},
		)
	}
}

func WithBaseTextHandler(level slog.Level, addSource bool) func(*LoggingHandler) {
	return func(handler *LoggingHandler) {
		handler.Handler = slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level:     level,
				AddSource: addSource,
			},
		)
	}
}

// Override the base Handle function to insert our attrs
func (handler *LoggingHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(slog.String("request-id", RetrieveStringFromCtx(ctx, RequestIdHeader, "unknown")))
	r.AddAttrs(slog.String("environment", handler.Environment))

	return handler.Handler.Handle(ctx, r)
}
