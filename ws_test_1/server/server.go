package server

import (
	snowflake_helper "code271/ws_test_1/pkg/snowflake"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ServerID string
	port     string
	engine   *gin.Engine
}

func (s *Server) NewServer(port string, r *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	s.port = port
	s.engine = r
	s.ServerID = fmt.Sprintf("server_%d", snowflake_helper.MakeID())
	fmt.Printf("服务：%s 开始启动。\n", s.ServerID)
}

func (s *Server) Run() (err error) {
	port := fmt.Sprintf(":%s", s.port)
	if err = s.engine.Run(port); err != nil {
		return
	}
	return
}
