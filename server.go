package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zhuliminl/easyrn-server/config"
	"github.com/zhuliminl/easyrn-server/controllers"
	"github.com/zhuliminl/easyrn-server/db"
	"github.com/zhuliminl/easyrn-server/docs"
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
	router.GET("/user/getUserByUserId", userController.GetUserByUserId)
	router.GET("/user/getMyInfo", userController.GetUserByUserIdBar)
	//
	//router.POST("/auth/registerByEmail", userController.GetUserById)
	//router.POST("/auth/registerByPhone", userController.GetUserById)
	//router.POST("/auth/loginByEmail", userController.GetUserById)
	//router.POST("/auth/loginByPhone", userController.GetUserById)
	//router.POST("/auth/wx/getOpenId", userController.GetUserById)
	//router.GET("/auth/wx/getMiniLink", userController.GetUserById)
	//router.GET("/auth/wx/getMiniLinkStatus", userController.GetUserById)
	//router.POST("/auth/wx/scanOver", userController.GetUserById)
	//router.POST("/auth/wx/loginWithEncryptedPhoneData", userController.GetUserById)

	// swagger 文档
	//docs.SwaggerInfo.Title = "easy react-native"
	//docs.SwaggerInfo.Description = "easyrn"
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "localhost"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 启动
	router.Run(address + ":" + port)
}
