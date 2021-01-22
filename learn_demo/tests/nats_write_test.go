package tests

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
)

func TestSendNats(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("nats 链接错误：", err.Error())
		return
	}
	name := "1349629906039869440"
	// 发布消息
	for i := 0; i < 5; i++ {
		err = nc.Publish(name, []byte("Hello f##k"))
		fmt.Println("发送成功！")
		if err != nil {
			t.Log("error:", err.Error())
		}
	}
	return
}
