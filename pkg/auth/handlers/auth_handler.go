package handlers

import (
	"ginfx-template/pkg/fx/ginfx"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	*jwt.GinJWTMiddleware
}

func (a *AuthHandler) Routes() []ginfx.Route {
	return []ginfx.Route{
		a.login(),
		a.refreshToken(),
	}
}

func NewAuthHandler(jwt *jwt.GinJWTMiddleware) *AuthHandler {
	return &AuthHandler{jwt}
}

var _ ginfx.Handler = (*AuthHandler)(nil)

func (a *AuthHandler) login() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "POST", "/login", a.LoginHandler
	}
}

func (a *AuthHandler) refreshToken() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "GET", "/refresh_token", a.RefreshHandler
	}
}
