package server

import (
	"fmt"
	"ginfx-template/config"
	"ginfx-template/internal"
	"ginfx-template/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
	"os"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}

func runServer(_ *cli.Context) error {
	hostname, _ := os.Hostname()
	slog.Info(fmt.Sprintf("Starting My App on %s with PID %d", hostname, os.Getpid()))
	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.SlogLogger{Logger: slog.Default()}
		}),
		fx.Provide(
			config.ProvideAppConfig,
			config.ProvideRedisClient,
			config.ProvideGormEngine,
			config.ProvideWeb,
		),
		fx.Options(auth.Module),
		fx.Options(internal.Module),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
	return nil
}
