package handler

import (
	"ginfx-template/internal/model"
	"ginfx-template/internal/service"
	"ginfx-template/pkg/fx/ginfx"
	"ginfx-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userSer service.IUserService) ginfx.Handler {
	return &UserHandler{
		userService: userSer,
	}
}

var _ ginfx.Handler = (*UserHandler)(nil)

func (u *UserHandler) Routes() []ginfx.Route {
	return []ginfx.Route{
		u.hello(),
		u.addUser(),
	}
}

func (u *UserHandler) hello() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "GET", "/api/users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
		}
	}
}

// @Summary      新增用户
// @Description
// @Tags         用户管理
// @Produce      json
// @Param        user  body  model.User  true  "User信息"
// @Success      200  {object}  string "ok"
// @Failure      400  {object}  string 报错信息
// @Router       /api/user/add [post]
func (u *UserHandler) addUser() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "POST", "/api/user/add", func(ctx *gin.Context) {
			var user model.User
			err := ctx.ShouldBind(&user)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "用户创建失败！请联系系统管理员！")
				return
			}
			err = u.userService.Save(&user)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "用户创建失败！请联系系统管理员!")
				return
			}
			ctx.JSON(http.StatusOK, "用户创建成功！")
		}
	}
}
