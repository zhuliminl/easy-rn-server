package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zhuliminl/easyrn-server/config"
	"github.com/zhuliminl/easyrn-server/controllers"
	"github.com/zhuliminl/easyrn-server/db"
	"github.com/zhuliminl/easyrn-server/docs"
	"github.com/zhuliminl/easyrn-server/middlewares"
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

	// redis 初始化
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer db.CloseDatabaseConnection()

	// 依赖注入
	var (
		userRepository repository.UserRepository = repository.NewUserRepository()
		userService    service.UserService       = service.NewUserService(userRepository)
		jwtService     service.JWTService        = service.NewJWTService()
		projectService service.ProjectService    = service.NewProjectService(userRepository)
		teamService    service.ProjectService    = service.NewTeamService(userRepository)

		wechatService service.WechatService = service.NewWechatService(userRepository, userService, rdb)
		authService   service.AuthService   = service.NewAuthService(userRepository, userService, jwtService)

		userController    controllers.UserController    = controllers.NewUserController(userService)
		projectController controllers.ProjectController = controllers.NewProjectController(userService, projectService)
		teamController    controllers.TeamController    = controllers.NewTeamController(userService, teamService)

		authController   controllers.AuthController   = controllers.NewAuthController(userService, authService, jwtService)
		wechatController controllers.WechatController = controllers.NewWechatController(wechatService, jwtService)
	)

	//  router 配置
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 中间件
	JWTMiddleware := middlewares.JWT(jwtService)

	// 路径配置
	// 用户
	router.GET("/user/getUserByUserId", userController.GetUserByUserId)
	router.GET("/user/getMyInfo", JWTMiddleware, userController.GetMyInfo)

	// 认证
	router.POST("/auth/registerByEmail", authController.RegisterByEmail)
	router.POST("/auth/registerByPhone", authController.RegisterByPhone)
	router.POST("/auth/loginByEmail", authController.LoginByEmail)
	router.POST("/auth/loginByPhone", authController.LoginByPhone)

	router.POST("/auth/wx/getOpenId", wechatController.GetOpenID)
	router.GET("/auth/wx/getMiniLink", wechatController.GetMiniLink)
	router.GET("/auth/wx/getMiniLinkStatus", wechatController.GetMiniLinkStatus)
	router.POST("/auth/wx/scanOver", wechatController.ScanOver)
	router.POST("/auth/wx/loginWithEncryptedPhoneData", wechatController.LoginWithEncryptedPhoneData)

	// 项目
	router.GET("/project/getAllProject", projectController.GetAllProject)
	router.POST("/project/createProject", projectController.CreateProject)
	router.POST("/project/updateProject", projectController.UpdateProject)
	router.POST("/project/deleteProject", projectController.DeleteProject)

	// 团队
	router.GET("/team/getAllTeam", teamController.GetAllTeam)
	router.POST("/team/createTeam", teamController.CreateTeam)
	router.POST("/team/updateTeam", teamController.UpdateTeam)
	router.POST("/team/deleteTeam", teamController.DeleteTeam)

	// swagger 文档
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "http://localhost:3500"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 启动
	router.Run(address + ":" + port)
}
