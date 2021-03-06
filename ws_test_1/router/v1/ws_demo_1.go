package v1

import (
	"code271/ws_test_1/pkg/server_context"
	"code271/ws_test_1/pkg/ws"
	service_1 "code271/ws_test_1/service/service_demo_1"
	"fmt"
	"net/http"
)

func WsPage(c *server_context.Context) {
	req := &struct {
		Key string `form:"key"`
	}{}
	if err := c.ShouldBindQuery(req); err != nil {
		fmt.Println("bind error: ", err.Error())
		http.NotFound(c.Writer, c.Request)
	}
	fmt.Println("the key: ", req.Key)
	if err := ws.RunWsClient(c.Context); err != nil {
		fmt.Println("run ws client error: ", err.Error())
		http.NotFound(c.Writer, c.Request)
	}
}

func Login(c *server_context.Context) {
	req := &struct {
		AccountName string `json:"account_name"`
		Password    string `json:"password"`
	}{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(server_context.NewSuccess(nil))
		return
	}
	code, data := service_1.Login(req.AccountName, req.Password)
	c.JSON(server_context.NewError(code, data))
	return
}

func Register(c *server_context.Context) {
	req := &struct {
		Mobile  string `json:"mobile"`
		SmsCode string `json:"sms_code"`
	}{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(server_context.NewSuccess(nil))
		return
	}
	code, data := service_1.Register(req.Mobile, req.SmsCode)
	c.JSON(server_context.NewError(code, data))
	return
}
