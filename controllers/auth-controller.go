package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constError"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/service"
)

type AuthController interface {
	RegisterByEmail(c *gin.Context)
	RegisterByPhone(c *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
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

	// 校验用户注册
	err = u.authService.VerifyRegisterByEmail(userRegister)
	if IsConstError(c, err, constError.EmailNotValid) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotValid) {
		return
	}
	if IsConstError(c, err, constError.UserDuplicated) {
		return
	}
	if Error500(c, err) {
		return
	}
	user, err := u.authService.CreateUserByEmail(userRegister)
	if Error500(c, err) {
		return
	}
	// 生成 token
	token := u.jwtService.GenerateToken(user.UserId)
	res := dto.ResRegister{Token: token, User: user}
	SendResponseOk(c, constant.RequestSuccess, res)
}

// RegisterByPhone
// @Tags      auth
// @Summary   通过手机注册
// @accept    application/json
// @Produce   application/json
// @Param	  data body dto.UserRegisterByPhone true "注册"
// @Success   200   {object}  Response{data=dto.ResRegister}  "用户信息和 token"
// @Router    /auth/registerByPhone [post]
func (u authController) RegisterByPhone(c *gin.Context) {
	var userRegister dto.UserRegisterByPhone
	err := c.ShouldBindJSON(&userRegister)
	if Error400(c, err) {
		return
	}

	// 校验用户注册
	err = u.authService.VerifyRegisterByPhone(userRegister)
	if IsConstError(c, err, constError.PhoneNumberNotValid) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotValid) {
		return
	}
	if IsConstError(c, err, constError.UserDuplicated) {
		return
	}
	if Error500(c, err) {
		return
	}
	user, err := u.authService.CreateUserByPhone(userRegister)
	if Error500(c, err) {
		return
	}
	// 生成 token
	token := u.jwtService.GenerateToken(user.UserId)
	res := dto.ResRegister{Token: token, User: user}
	SendResponseOk(c, constant.RequestSuccess, res)
}

func NewAuthController(userService service.UserService, authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{userService: userService, authService: authService, jwtService: jwtService}
}
