package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/config"
	"github.com/zhuliminl/easyrn-server/controllers"
	"github.com/zhuliminl/easyrn-server/db"
	"github.com/zhuliminl/easyrn-server/repository"
	"github.com/zhuliminl/easyrn-server/service"
)

func StartServer() {

	// 读取配置
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	// 数据库启动
	db.Init()

	defer db.CloseDatabaseConnection()

	// 依赖注入
	var (
		userRepository repository.UserRepository  = repository.NewUserRepository()
		userService    service.UserService        = service.NewUserService(userRepository)
		userController controllers.UserController = controllers.NewUserController(userService)
	)
	//  router 配置
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 路径配置
	router.GET("/user/getUserByUserId", userController.GetUserById)
	router.GET("/user/getMyInfo", userController.GetUserById)

	router.POST("/auth/registerByEmail", userController.GetUserById)
	router.POST("/auth/registerByPhone", userController.GetUserById)
	router.POST("/auth/loginByEmail", userController.GetUserById)
	router.POST("/auth/loginByPhone", userController.GetUserById)
	router.POST("/auth/wx/getOpenId", userController.GetUserById)
	router.GET("/auth/wx/getMiniLink", userController.GetUserById)
	router.GET("/auth/wx/getMiniLinkStatus", userController.GetUserById)
	router.POST("/auth/wx/scanOver", userController.GetUserById)
	router.POST("/auth/wx/loginWithEncryptedPhoneData", userController.GetUserById)

	// 启动
	router.Run(address + ":" + port)
}
