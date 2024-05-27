package logger

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

type gormLogger struct {
}

var _ logger.Interface = (*gormLogger)(nil)

func NewGormLogger() logger.Interface {
	return &gormLogger{}
}

func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	slog.InfoContext(ctx, s, fmt.Sprintf(s, i))
}

func (g *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	slog.WarnContext(ctx, s, fmt.Sprintf(s, i))
}

func (g *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	slog.ErrorContext(ctx, s, fmt.Sprintf(s, i))
}

func (g *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {

}
