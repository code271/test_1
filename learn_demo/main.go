package main

import (
	"code271/learn_demo/pkg/nats_helper"
	"code271/learn_demo/pkg/ws"
	"code271/learn_demo/router"
	"code271/learn_demo/server"
	"fmt"
)

func main() {
	if err := nats_helper.NatsInit(); err != nil {
		fmt.Println("nats初始化失败：", err.Error())
		return
	}
	go ws.Manager.Start()
	r := router.SetRouter()
	s := new(server.Server)
	s.NewServer("9050", r)
	if err := s.Run(); err != nil {
		fmt.Println("服务启动错误：", err.Error())
		return
	}
	return
}
