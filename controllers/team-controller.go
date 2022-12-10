package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/service"
)

type TeamController interface {
	GetAllTeam(c *gin.Context)
	CreateTeam(c *gin.Context)
	UpdateTeam(c *gin.Context)
	DeleteTeam(c *gin.Context)
}

type teamController struct {
	userService service.UserService
	teamService service.TeamService
}

func (t teamController) GetAllTeam(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t teamController) CreateTeam(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t teamController) UpdateTeam(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t teamController) DeleteTeam(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewTeamController(userService service.UserService, teamService service.TeamService) TeamController {
	return &teamController{userService: userService, teamService: teamService}
}
