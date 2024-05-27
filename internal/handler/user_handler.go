package handler

import (
	"ginfx-template/pkg/fx/ginfx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() ginfx.Handler {
	return &UserHandler{}
}

func (u *UserHandler) Routes() []ginfx.Route {
	return []ginfx.Route{
		u.hello(),
	}
}

func (u *UserHandler) hello() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "GET", "/api/users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
		}
	}
}

var _ ginfx.Handler = (*UserHandler)(nil)
