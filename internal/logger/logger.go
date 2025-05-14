package logger

import (
	"log/slog"
	"os"
)

// New создает новый экземпляр slog логгера, в зависимости от уровня окружения.
// В случае окружения RROD логгер будт выводить информацию на уровне info и выше.
// На другом уровне оркжения - начиная с уровня dubug.
func New(env string) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if env == "PROD" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return logger

}
