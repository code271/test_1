package test_1_test

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
)

var ncConn *nats.Conn

func TestListenNats(t *testing.T) {
	c := new(client)
	var err error
	ncConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("链接nats失败：", err.Error())
		return
	}
	c.Send = make(chan []byte, 1000000)
	c.Count = make(chan int, 1000000)
	sub, err := c.ListenNats("foo")
	if err != nil {
		t.Log("监听失败：", err.Error())
		return
	}
	defer func() {
		err = sub.Unsubscribe()
		if err != nil {
			t.Log("退订还能失败：", err.Error())
		}
	}()
	if err != nil {
		t.Log("链接nats失败：", err.Error())
		return
	}
	for {
		select {
		case a, ok := <-c.Send:
			if !ok {
				fmt.Println("管道被关闭，结束请求")
				return
			}
			go func() {
				theCount := <-c.Count
				fmt.Println(string(a), "这是第：", theCount, " 条数据。")
				if theCount%500000 == 0 {
					fmt.Println("asdfasdfasdfasdfasdfasdf")
				}
			}()
		}
	}
}

func (c *client) ListenNats(name string) (sub *nats.Subscription, err error) {
	//name = "foo"
	fmt.Println("开始监听：", name)
	sub, err = ncConn.Subscribe(name, func(m *nats.Msg) {
		//go fmt.Println("收到消息了：", string(m.Data))
		c.CountFlat++
		c.Count <- c.CountFlat
		c.Send <- m.Data
	})
	if err != nil {
		fmt.Println("nats listen error：", err.Error())
		return
	}
	return
}

type client struct {
	Send      chan []byte
	Count     chan int
	CountFlat int
}
