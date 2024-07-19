package ginfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler interface {
	Routes() []Route
}

func AsAuthHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"auth_handlers"`),
	)
}

func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

type Route func() (method string, pattern string, handler gin.HandlerFunc)

type Handlers struct {
	fx.In
	Handlers     []Handler `group:"handlers"`
	AuthHandlers []Handler `group:"auth_handlers"`
}
