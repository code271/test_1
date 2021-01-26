package test_5_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestIt(t *testing.T) {
	r := gin.Default()
	r.GET("/t", func(c *gin.Context) {
		// 外部重定向
		// 直接跳转界面。。。一般用不上
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/move", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		// 内部重定向
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	r.Run()
}
