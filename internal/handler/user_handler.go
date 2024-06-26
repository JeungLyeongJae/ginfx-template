package handler

import (
	"ginfx-template/internal/model"
	"ginfx-template/internal/model/response"
	"ginfx-template/internal/service"
	"ginfx-template/pkg/fx/ginfx"
	"ginfx-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		u.isExistByUsername(),
		u.getUserList(),
		u.addUser(),
		u.updateUser(),
		u.deleteUser(),
	}
}

// IsExistByUsername
// @Summary      用户列表
// @Description
// @Tags         用户管理
// @Produce      json
// @Param        username {object} string true  "查询用户名"
// @Success      200  {object}  boolean 是否存在
// @Failure      400  {object}  string 报错信息
// @Router       /api/user/is_exist_username [get]
func (u *UserHandler) isExistByUsername() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "GET", "/api/user/is_exist_username", func(ctx *gin.Context) {
			username := ctx.Query("username")
			isExist, err := u.userService.IsExistByUsername(username)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "系统查询失败！请联系系统管理员!")
				return
			}
			ctx.JSON(http.StatusOK, isExist)
		}
	}
}

// @Summary      用户列表
// @Description
// @Tags         用户管理
// @Produce      json
// @Param        user  body  response.Page  true  "分页信息"
// @Success      200  body  response.Page 分页数据
// @Failure      400  {object}  string 报错信息
// @Router       /api/user/list [get]
func (u *UserHandler) getUserList() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "GET", "/api/user/list", func(ctx *gin.Context) {
			var page response.Page

			page.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
			page.PageNumber, _ = strconv.Atoi(ctx.Query("page_number"))
			page.Condition = ctx.Query("condition")

			err := u.userService.GetUserList(&page)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "系统查询失败！请联系系统管理员!")
				return
			}
			ctx.JSON(http.StatusOK, page)
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
			err = u.userService.Create(&user)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "用户创建失败！请联系系统管理员!")
				return
			}
			ctx.JSON(http.StatusOK, "用户创建成功！")
		}
	}
}

// UpdateUser
// @Summary      编辑用户
// @Description
// @Tags         用户管理
// @Produce      json
// @Param        user  body  model.User  true  "User信息"
// @Success      200  {object}  string "ok"
// @Failure      400  {object}  string 报错信息
// @Router       /api/user/update [post]
func (u *UserHandler) updateUser() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "POST", "/api/user/update", func(ctx *gin.Context) {
			var user model.User
			err := ctx.ShouldBind(&user)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "用户保存失败！请联系系统管理员！")
				return
			}
			err = u.userService.Update(&user)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "用户保存失败！请联系系统管理员!")
				return
			}
			ctx.JSON(http.StatusOK, "用户保存成功！")
		}
	}
}

// @Summary      删除用户
// @Description
// @Tags         用户管理
// @Produce      json
// @Param        user.id  {object}  string  true  "UserId"
// @Success      200  {object}  string "ok"
// @Failure      400  {object}  string 报错信息
// @Router       /api/user/delete [get]
func (u *UserHandler) deleteUser() ginfx.Route {
	return func() (method string, pattern string, handler gin.HandlerFunc) {
		return "GET", "/api/user/delete", func(ctx *gin.Context) {
			var user model.User
			user.Id, _ = strconv.ParseInt(ctx.Query("id"), 0, 64)
			err := u.userService.Delete(&user)
			if err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusBadRequest, "用户删除失败！请联系系统管理员!")
				return
			}
			ctx.JSON(http.StatusOK, "用户删除成功！")
		}
	}
}
