package routes

import (
	"admin/app/http/controller"
	"admin/app/http/controller/system"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	indexGroup := r.Group("/")
	{
		//首页
		indexGroup.GET("/", controller.Index)
	}

	systemAccountGroup := r.Group("account")
	{
		//增加用户User
		systemAccountGroup.GET("/create", system.CreateUser)
	}

	return r
}
