package router

import (
	"github.com/mouuii/go-easy-blog/api/user/handler"

	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	engine := gin.Default()
	wx := engine.Group("user")
	
	InitRouterV1(wx.Group("v1"))

	return engine
}


func InitRouterV1(v1 *gin.RouterGroup) {

	
	v1.POST("/login", handler.Login)
	
}