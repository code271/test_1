package main

import (
	"code271/ws_test_1/pkg/const_key"
	"code271/ws_test_1/pkg/db"
	"code271/ws_test_1/pkg/nats_helper"
	snowflake_helper "code271/ws_test_1/pkg/snowflake"
	"code271/ws_test_1/pkg/ws"
	"code271/ws_test_1/router"
	"code271/ws_test_1/server"
	"fmt"
)

func main() {
	if err := nats_helper.NatsInit(); err != nil {
		fmt.Println("nats初始化失败：", err.Error())
		return
	}

	{
		if err := db.Init(func(config *db.Config) {
			config.DataSourceName = const_key.SqlUrl
			config.LogMode = true
			config.MaxIdleConn = 100
			config.MaxOpenConn = 300
			config.MaxLifetime = 0
			config.PrefixMapper = "t_"
		}); err != nil {
			panic(err)
		}
		defer db.Close()
	}

	_ = snowflake_helper.Init()
	r := router.SetRouter()
	s := new(server.Server)
	s.NewServer("9050", r)
	go ws.Center.Start(s.ServerID)
	if err := s.Run(); err != nil {
		fmt.Println("服务启动错误：", err.Error())
		return
	}
	return
}
