package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/service"
)

type UserController interface {
	GetUserById(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (u userController) GetUserById(c *gin.Context) {
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
