package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/dto"
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

// CreateTeam
// @Tags      team
// @Summary   创建一个团队
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  Response{data=dto.TeamCreate}  "团队基础信息"
// @Router    /team/createTeam [post]
func (t teamController) CreateTeam(c *gin.Context) {
	userId := c.MustGet("CurrentUserId").(string)
	if userId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}

	var teamCreate dto.TeamCreate
	err := c.ShouldBindJSON(&teamCreate)
	if Error400(c, err) {
		return
	}

	err = t.teamService.CreateTeam(userId, teamCreate)
	if Error500(c, err) {
		return
	}

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
