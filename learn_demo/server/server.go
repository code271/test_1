package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func (s *Server) NewServer(port string, r *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	s.port = port
	s.engine = r
}

func (s *Server) Run() (err error) {
	port := fmt.Sprintf(":%s", s.port)
	if err = s.engine.Run(port); err != nil {
		return
	}
	return
}
