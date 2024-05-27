package config

import (
	"context"
	"fmt"
	"ginfx-template/pkg/fx/ginfx"
	"ginfx-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
)

func ProvideWeb(lc fx.Lifecycle, handlers ginfx.Handlers, appConfig *AppConfig) *gin.Engine {
	if strings.ToLower(appConfig.ServerConfig.Mode) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	app.Use(gin.Recovery())
	app.NoRoute(func(context *gin.Context) {
		if context.Request.RequestURI != "/" {
			context.Redirect(http.StatusTemporaryRedirect, "/")
		}
	})
	for _, handler := range handlers.Handlers {
		for _, route := range handler.Routes() {
			app.Handle(route())
		}
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.ServerConfig.Port),
		Handler: app,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			slog.Info("Web server started on port(s)" + srv.Addr)
			slog.Info("Press Ctrl+C to stop")
			go func() {
				if err0 := srv.Serve(ln); err0 != nil {
					slog.Error(err0.Error())
				}
			}()
			return nil
		},
		OnStop: srv.Shutdown,
	})
	return app
}

func ProvideGormEngine(appConfig *AppConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(appConfig.DbConfig.dsn()), &gorm.Config{
		Logger: logger.NewGormLogger(),
	})
	if err != nil {
		slog.Error(fmt.Sprint(err))
		os.Exit(1)
	}
	return db
}

func ProvideRedisClient(config *AppConfig) redis.UniversalClient {
	var rdb redis.UniversalClient

	switch config.RedisConfig.Addr {
	case "cluster":
		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{config.RedisConfig.Addr},
			Password: config.RedisConfig.Password,
		})
	case "sentinel":
		rdb = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    config.RedisConfig.MasterName,
			SentinelAddrs: []string{config.RedisConfig.Addr},
			Password:      config.RedisConfig.Password,
		})
	default:
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.RedisConfig.Addr,
			Password: config.RedisConfig.Password,
			DB:       config.RedisConfig.DB,
		})
	}

	return rdb
}
