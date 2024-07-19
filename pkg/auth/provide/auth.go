package provide

import (
	"errors"
	"ginfx-template/pkg/auth/domain"
	"ginfx-template/pkg/auth/services"
	"ginfx-template/pkg/common"
	"log"
	"log/slog"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	identityKey = "sub"
)

func NewGinJWTMiddleware(userService services.UserDetailsService) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "model-builder",
		SigningAlgorithm: "HS256",
		Key:              []byte("gin fx template"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(domain.UserDetails); ok {
				return jwt.MapClaims{
					"iss":         "model-builder",
					identityKey:   v.GetUsername(),
					"name":        v.GetName(),
					"authorities": v.GetAuthorities(),
					"last_login":  v.GetLastLogin().Format(time.RFC3339),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &domain.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginRequest struct {
				Username string `json:"username" binding:"required"`
				Password string `json:"password" binding:"required"`
			}
			if err := c.ShouldBindJSON(&loginRequest); err != nil {
				slog.Error(err.Error())
				return "", jwt.ErrMissingLoginValues
			}

			userDetails, err := userService.LoadUserByUsername(loginRequest.Username)
			if err != nil {
				return nil, errors.New("登录失败")
			}

			if common.Encoder.Matches(loginRequest.Password, userDetails.GetPassword()) {
				err := userService.UpdateLastLogin(loginRequest.Username)
				if err != nil {
					return nil, err
				} // 更新最后登录时间
				return userDetails, nil
			}
			return nil, errors.New("登录失败")
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(domain.UserDetails); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

func GetUserName(c *gin.Context) string {
	u, exist := c.Get(identityKey)
	if exist {
		return u.(*domain.User).Username
	}
	return ""
}
