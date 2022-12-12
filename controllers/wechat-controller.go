package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constError"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/service"
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
// @Param	  login_session_id query string true "login_session_id"
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
	if statusDto.Status == strconv.FormatInt(constant.WechatLoginSuccess, 10) {
		userDto, err := u.wechatService.GetUserByLoginSessionId(loginSessionId)
		if Error500(c, err) {
			return
		}

		token := u.jwtService.GenerateToken(userDto.ID)
		res := dto.ResRegister{Token: token, User: userDto}
		SendResponseOk(c, constant.RequestSuccess, res)
		return
	}
	SendResponseOk(c, constant.RequestSuccess, statusDto)
}

// ScanOver
// @Tags      auth
// @Summary   记录小程序被扫描
// @accept    application/json
// @Produce   application/json
// @Param	    data body dto.LinkScanOver true "xx"
// @Router    /auth/wx/scanOver [post]
func (u wechatController) ScanOver(c *gin.Context) {
	var scan dto.LinkScanOver
	err := c.ShouldBindJSON(&scan)
	if Error400(c, err) {
		return
	}

	err = u.wechatService.ScanOver(scan.LoginSessionId)
	if IsConstError(c, err, constError.WechatLoginUidNotFound) {
		return
	}

	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, EmptyObj{})
}

// LoginWithEncryptedPhoneData
// @Tags      auth
// @Summary   微信通过加密数据登录
// @accept    application/json
// @Produce   application/json
// @Param	    data body dto.WxLoginData true "WxLoginData"
// @Router    /auth/wx/loginWithEncryptedPhoneData [post]
func (u wechatController) LoginWithEncryptedPhoneData(c *gin.Context) {
	var wxLoginData dto.WxLoginData
	err := c.ShouldBindJSON(&wxLoginData)
	if Error400(c, err) {
		return
	}

	resWxLogin, err := u.wechatService.LoginWithEncryptedPhoneData(wxLoginData)
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, resWxLogin)
}

// GetOpenID
// @Tags      auth
// @Summary   获取 openId
// @accept    application/json
// @Produce   application/json
// @Param	    data body dto.WechatCodeDto true "WechatCodeDto"
// @Router    /auth/wx/getOpenId [post]
func (u wechatController) GetOpenID(c *gin.Context) {
	var wechatCode dto.WechatCodeDto
	err := c.ShouldBindJSON(&wechatCode)
	if Error400(c, err) {
		return
	}

	session, err := u.wechatService.GetOpenId(wechatCode)
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, session)
}

func NewWechatController(wechatService service.WechatService, jwtService service.JWTService) WechatController {
	return &wechatController{wechatService: wechatService, jwtService: jwtService}
}
