package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/service"
)

type UserController interface {
	GetUserByUserId(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

// GetUserByUserId
// @Tags      user
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  userId query string true "用户 userId"
// @Success   200   {object}  Response{data=dto.User}  "用户信息"
// @Router    /user/getUserByUserId [get]
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

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}
