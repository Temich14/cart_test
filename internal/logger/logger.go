package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if env == "PROD" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return logger

}
