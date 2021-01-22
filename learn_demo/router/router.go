package router

import (
	"code271/learn_demo/pkg/middleware"
	v1 "code271/learn_demo/router/v1"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	{ //demo 1
		wsApi := router.Group("/v1/demo1")
		wsApi.GET("/wsDemo1", v1.WsPage)
	}
	router.Use(middleware.Cors())
	return router
}
