package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if env == "PROD" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return logger

}
