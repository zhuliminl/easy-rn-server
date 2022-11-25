package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/easyrn-server/config"
)

func StartServer() {

	// 读取配置
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	//  router 配置
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 启动
	router.Run(address + ":" + port)
}
