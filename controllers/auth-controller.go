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
	LoginByEmail(c *gin.Context)
	LoginByPhone(c *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
}

// LoginByEmail
// @Tags      auth
// @Summary   通过邮箱登录
// @accept    application/json
// @Produce   application/json
// @Param	  data body dto.UserLoginByEmail true "登录"
// @Success   200   {object}  Response{data=dto.ResToken}  "用户信息和 token"
// @Router    /auth/loginByEmail [post]
func (u authController) LoginByEmail(c *gin.Context) {
	var userLogin dto.UserLoginByEmail
	err := c.ShouldBindJSON(&userLogin)
	if Error400(c, err) {
		return
	}
	user, err := u.authService.VerifyCredentialByEmail(userLogin.Email, userLogin.Password)
	if IsConstError(c, err, constError.UserNotFound) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotMatch) {
		return
	}
	if Error500(c, err) {
		return
	}

	token := u.jwtService.GenerateToken(user.ID)
	SendResponseOk(c, constant.LoginSuccess, dto.ResToken{Token: token})
}

// LoginByPhone
// @Tags      auth
// @Summary   通过手机号登录
// @accept    application/json
// @Produce   application/json
// @Param	  data body dto.UserLoginByPhone true "登录"
// @Success   200   {object}  Response{data=dto.ResRegister}  "用户信息和 token"
// @Router    /auth/loginByPhone [post]
func (u authController) LoginByPhone(c *gin.Context) {
	var userLogin dto.UserLoginByPhone
	err := c.ShouldBindJSON(&userLogin)
	if Error400(c, err) {
		return
	}
	user, err := u.authService.VerifyCredentialByPhone(userLogin.Phone, userLogin.Password)
	if IsConstError(c, err, constError.UserNotFound) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotMatch) {
		return
	}
	if Error500(c, err) {
		return
	}

	token := u.jwtService.GenerateToken(user.ID)
	SendResponseOk(c, constant.LoginSuccess, dto.ResToken{Token: token})
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
	token := u.jwtService.GenerateToken(user.ID)
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
	token := u.jwtService.GenerateToken(user.ID)
	res := dto.ResRegister{Token: token, User: user}
	SendResponseOk(c, constant.RequestSuccess, res)
}

func NewAuthController(userService service.UserService, authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{userService: userService, authService: authService, jwtService: jwtService}
}
