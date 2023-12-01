package apputil

import (
	"log/slog"
	"os"
)

func NewSlogDefault(level slog.Level) {
	handler := slog.NewJSONHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		},
	)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func SetTraceId(traceId string) {
	slog.Default().With("trace_id", traceId)
}

func SetCorrelationId(correlationId string) {
	slog.Default().With("correlation_id", correlationId)
}
