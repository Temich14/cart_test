package logger

import (
	"context"
	"errors"
	"log/slog"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

// GormLogger обертка над slog для взаимодействия с gorm
type GormLogger struct {
	logger    *slog.Logger
	level     gormlogger.LogLevel
	slowQuery time.Duration
}

func NewGormLogger(logger *slog.Logger, level gormlogger.LogLevel) *GormLogger {
	return &GormLogger{
		logger:    logger,
		level:     level,
		slowQuery: 200 * time.Millisecond,
	}
}

func (g *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	g.level = level
	return g
}

func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if g.level >= gormlogger.Info {
		g.logger.InfoContext(ctx, msg, slog.Any("data", data))
	}
}

func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if g.level >= gormlogger.Warn {
		g.logger.WarnContext(ctx, msg, slog.Any("data", data))
	}
}

func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if g.level >= gormlogger.Error {
		g.logger.ErrorContext(ctx, msg, slog.Any("data", data))
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.level == gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && !errors.Is(err, gormlogger.ErrRecordNotFound):
		g.logger.ErrorContext(ctx, "db error",
			slog.String("error", err.Error()),
			slog.String("sql", sql),
			slog.Int64("rows", rows),
			slog.Duration("elapsed", elapsed),
		)
	case elapsed > g.slowQuery:
		g.logger.WarnContext(ctx, "slow query",
			slog.String("sql", sql),
			slog.Int64("rows", rows),
			slog.Duration("elapsed", elapsed),
		)
	default:
		if g.level >= gormlogger.Info {
			g.logger.DebugContext(ctx, "query executed",
				slog.String("sql", sql),
				slog.Int64("rows", rows),
				slog.Duration("elapsed", elapsed),
			)
		}
	}
}
