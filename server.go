package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/config"
	"github.com/zhuliminl/easyrn-server/database"
	"gorm.io/gorm"
)

func StartServer() {

	// 读取配置
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	// 数据库启动
	var (
		db *gorm.DB = database.SetupDatabaseConnection()
	)
	defer database.CloseDatabaseConnection(db)

	//  router 配置
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 启动
	router.Run(address + ":" + port)
}
