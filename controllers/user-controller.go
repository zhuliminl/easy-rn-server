package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/service"
)

type UserController interface {
	GetUserByUserId(c *gin.Context)
	GetUserByUserIdBar(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

// GetUserByUserId
// PingExample godoc
// @Summary	sum
// @Schemes
// @Description	    获取用户信息
// @Tags			user
// @Accept			json
// @Produce		    json
// @Param           userId     query     string     false  "用户Id" xx
// @Param           id     body     dto.User     true  "用户Id" xx
// @Success	        200	{object}	dto.User
// @Router			/user/getUserByUserId [get]
func (u userController) GetUserByUserId(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	user, err := u.userService.GetUserByUserId(userId)
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, user)
}

// GetUserByUserIdBar
// PingExample godoc
// @Summary	sum
// @Schemes
// @Description	    获取用户信息
// @Tags			user
// @Accept			json
// @Produce		    json
// @Param           userId     query     string     false  "用户Id" xx
// @Success	        200	{object}	dto.User
// @Router			/user/getUserByUserIdBar [get]
func (u userController) GetUserByUserIdBar(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	user, err := u.userService.GetUserByUserId(userId)
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, user)
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}
