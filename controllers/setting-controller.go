package controllers

import (
	"github.com/zhuliminl/easyrn-server/service"
)

type SettingController interface {
}

type settingController struct {
	userService    service.UserService
	settingService service.SettingService
}

func NewSettingController(userService service.UserService, settingService service.SettingService) SettingController {
	return &teamController{userService: userService, teamService: settingService}
}
