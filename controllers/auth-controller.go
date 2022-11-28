package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/service"
)

type AuthController interface {
	RegisterByEmail(c *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
}

// RegisterByEmail
// @Tags      auth
// @Summary   通过邮箱注册
// @accept    application/json
// @Produce   application/json
// @Param	  data body dto.UserRegisterByEmail true "注册"
// @Success   200   {object}  Response{data=dto.ResRegister}  "用户信息和 token"
// @Router    /auth/registerByEmail [post]
func (u authController) RegisterByEmail(c *gin.Context) {
	var userRegister dto.UserRegisterByEmail
	err := c.ShouldBindJSON(&userRegister)
	if Error400(c, err) {
		return
	}

	//err = u.userService.

}

func NewAuthController(userService service.UserService, authService service.AuthService) AuthController {
	return &authController{userService: userService, authService: authService}
}
