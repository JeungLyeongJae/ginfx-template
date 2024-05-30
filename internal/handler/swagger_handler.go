package handler

import (
	"ginfx-template/pkg/fx/ginfx"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerHandler struct {
}

func NewSwaggerHandler() ginfx.Handler {
	return &SwaggerHandler{}
}

var _ ginfx.Handler = (*SwaggerHandler)(nil)

func (u *SwaggerHandler) Routes() []ginfx.Route {
	if gin.Mode() == "debug" {
		return []ginfx.Route{
			func() (method string, pattern string, handler gin.HandlerFunc) {
				return "GET", "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
					c.InstanceName = "swagger"
				})
			},
		}
	}
	return []ginfx.Route{}
}
