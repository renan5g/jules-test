package logging

import (
	"context"
	"log/slog"
	"os"
	"github.com/generic-org/payment-gateway/pkg/application/port"
)

type SlogAdapter struct {
	logger *slog.Logger
}

// NewSlogAdapter creates a new SlogAdapter.
// Level can be "debug", "info", "warn", "error".
func NewSlogAdapter(level string, isJSON bool) port.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
		// AddSource: true, // Uncomment to include source file and line number
	}

	var handler slog.Handler
	if isJSON {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	
	logger := slog.New(handler)
	return &SlogAdapter{logger: logger}
}

func (s *SlogAdapter) Debug(ctx context.Context, msg string, args ...interface{}) {
	s.logger.DebugContext(ctx, msg, args...)
}

func (s *SlogAdapter) Info(ctx context.Context, msg string, args ...interface{}) {
	s.logger.InfoContext(ctx, msg, args...)
}

func (s *SlogAdapter) Warn(ctx context.Context, msg string, args ...interface{}) {
	s.logger.WarnContext(ctx, msg, args...)
}

func (s *SlogAdapter) Error(ctx context.Context, msg string, args ...interface{}) {
	s.logger.ErrorContext(ctx, msg, args...)
}

func (s *SlogAdapter) Fatal(ctx context.Context, msg string, args ...interface{}) {
	s.logger.ErrorContext(ctx, msg, args...) // slog doesn't have Fatal, so log as error
	os.Exit(1)
}
