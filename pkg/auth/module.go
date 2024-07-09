package auth

import (
	"ginfx-template/pkg/auth/handlers"
	"ginfx-template/pkg/auth/provide"
	"ginfx-template/pkg/auth/services"
	"ginfx-template/pkg/fx/ginfx"
	"go.uber.org/fx"
)

var Module = fx.Module("AuthModule",
	fx.Provide(
		services.NewUserDetailsService,
	),
	fx.Provide(
		provide.NewGinJWTMiddleware,
		provide.NewEnforcer,
		provide.NewAuthorizer,
	),
	fx.Provide(
		ginfx.AsHandler(handlers.NewAuthHandler),
	),
)
