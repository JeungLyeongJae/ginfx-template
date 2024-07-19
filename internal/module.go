package internal

import (
	"ginfx-template/internal/handler"
	"ginfx-template/internal/model"
	"ginfx-template/internal/repository"
	"ginfx-template/internal/service"
	"ginfx-template/pkg/fx/ginfx"
	"ginfx-template/pkg/logger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("App",
	fx.Options(
		fx.Provide(repository.NewUserRepo),
	),
	fx.Options(
		fx.Provide(service.NewUserService),
	),
	fx.Options(
		fx.Provide(ginfx.AsAuthHandler(handler.NewUserHandler)),
		fx.Provide(ginfx.AsHandler(handler.NewSwaggerHandler)),
	),
	fx.Invoke(AutoMigration),
)

func AutoMigration(db *gorm.DB) {
	// 自动迁移
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		logger.Error("failed to migrate database")
		return
	}
	logger.Info("Database migration completed successfully")
}
