package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/service"
)

type ProjectController interface {
	GetAllProject(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type projectController struct {
	userService    service.UserService
	projectService service.ProjectService
}

func (p projectController) GetAllProject(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) CreateProject(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) UpdateProject(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DeleteProject(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewProjectController(userService service.UserService, projectService service.ProjectService) ProjectController {
	return &projectController{userService: userService, projectService: projectService}
}
