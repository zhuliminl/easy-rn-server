package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	GetUserById(context *gin.Context)
}

type userController struct {
	// userService
}
