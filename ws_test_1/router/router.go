package router

import (
	"code271/ws_test_1/pkg/middleware"
	"code271/ws_test_1/pkg/server_context"
	v1 "code271/ws_test_1/router/v1"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	{ //demo 1
		wsApi := router.Group("/v1/demo1")
		wsApi.GET("/wsDemo1", server_context.Handle(v1.WsPage))
		wsApi.POST("/login", server_context.Handle(v1.Login))
		wsApi.POST("/register", server_context.Handle(v1.Register))

	}
	router.Use(middleware.Cors())
	return router
}
