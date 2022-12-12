package controllers

import (
	"errors"

	"github.com/zhuliminl/easyrn-server/constError"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constant"
	_ "github.com/zhuliminl/easyrn-server/dto"
	_ "github.com/zhuliminl/easyrn-server/entity"
	"github.com/zhuliminl/easyrn-server/service"
)

type UserController interface {
	GetUserByUserId(c *gin.Context)
	GetMyInfo(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

// GetMyInfo
// @Tags      user
// @Summary   获取自己的用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  Response{data=entity.User}  "用户信息"
// @Router    /user/getMyInfo [get]
func (u userController) GetMyInfo(c *gin.Context) {
	userId := c.MustGet("CurrentUserId").(string)
	if userId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	user, err := u.userService.GetUserByUserId(userId)
	if IsConstError(c, err, constError.UserNotFound) {
		return
	}
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, user)
}

// GetUserByUserId
// @Tags      user
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  userId query string true "用户 userId"
// @Success   200   {object}  Response{data=entity.User}  "用户信息"
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
