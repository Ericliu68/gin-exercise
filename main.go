package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"gin-user/api/login"
	"gin-user/db"
	"gin-user/settings"
	"gin-user/utils"
)

func main() {
	// 初始化配置
	utils.LogerInit()
	utils.CreatePathPemKey()
	settings.GetConfig()
	db.DbInit()
	router := gin.Default()
	pprof.Register(router)
	router.Use(cors.Default())
	v1 := router.Group("/v1/login_service")
	{
		v1.POST("/login", login.Login)
	}

	router.Run(":8000")
}
