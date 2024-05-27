package internal

import (
	"ginfx-template/internal/handler"
	"ginfx-template/internal/service"
	"ginfx-template/pkg/fx/ginfx"
	"go.uber.org/fx"
)

var Module = fx.Module("App",
	fx.Options(
		fx.Provide(service.NewUserService),
	),
	fx.Options(
		fx.Provide(ginfx.AsHandler(handler.NewUserHandler)),
	),
)
