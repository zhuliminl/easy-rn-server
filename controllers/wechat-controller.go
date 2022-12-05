package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constError"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/service"
	"strconv"
)

type WechatController interface {
	GetOpenID(c *gin.Context)
	GetMiniLink(c *gin.Context)
	GetMiniLinkStatus(c *gin.Context)
	ScanOver(c *gin.Context)
	LoginWithEncryptedPhoneData(c *gin.Context)
}

type wechatController struct {
	wechatService service.WechatService
	jwtService    service.JWTService
}

// GetMiniLink
// @Tags      auth
// @Summary   获取小程序跳转链接
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  Response{data=dto.WechatAppLink}  "小程序链接"
// @Router    /auth/wx/getMiniLink [get]
func (u wechatController) GetMiniLink(c *gin.Context) {
	linkDto, err := u.wechatService.GenerateAppLink()
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, linkDto)
}

// GetMiniLinkStatus
// @Tags      auth
// @Summary   web 端轮询当前登录链接状态
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  Response{data=dto.MiniLinkStatus}  "小程序链接"
// @Router    /auth/wx/getMiniLinkStatus [get]
func (u wechatController) GetMiniLinkStatus(c *gin.Context) {
	loginSessionId := c.Query("login_session_id")
	if loginSessionId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	statusDto, err := u.wechatService.GetMiniLinkStatus(loginSessionId)
	if IsConstError(c, err, constError.WechatLoginUidNotFound) {
		return
	}
	if Error500(c, err) {
		return
	}
	if statusDto.Status == strconv.FormatInt(0, 10) {
		userDto, err := u.wechatService.GetUserByLoginSessionId(loginSessionId)
		if Error500(c, err) {
			return
		}

		token := u.jwtService.GenerateToken(userDto.UserId)
		res := dto.ResRegister{Token: token, User: userDto}
		SendResponseOk(c, constant.RequestSuccess, res)
		return
	}
	SendResponseOk(c, constant.RequestSuccess, statusDto)
}

func (u wechatController) ScanOver(c *gin.Context) {
}

func (u wechatController) LoginWithEncryptedPhoneData(c *gin.Context) {
}

// GetOpenID
// @Tags      auth
// @Summary   获取 openId
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  Response{data=dto.User}  "用户信息"
// @Router    /auth/wx/getOpenId [post]
func (u wechatController) GetOpenID(c *gin.Context) {
}

func NewWechatController(wechatService service.WechatService, jwtService service.JWTService) WechatController {
	return &wechatController{wechatService: wechatService, jwtService: jwtService}
}
