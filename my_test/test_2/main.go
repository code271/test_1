package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

var ncConn *nats.Conn

func main() {
	c := new(client)
	var err error
	ncConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("链接nats失败：", err.Error())
		return
	}
	c.Send = make(chan []byte, 1000000)
	c.Count = make(chan int, 1000000)
	err = c.ListenNats("foo")
	if err != nil {
		fmt.Println("链接nats失败：", err.Error())
		return
	}
	for {
		select {
		case a, ok := <-c.Send:
			if !ok {
				return
			}
			//go func() {
				theCount := <-c.Count
				fmt.Println(string(a), "这是第：", theCount, " 条数据")
				if theCount%500000 == 0 {
					fmt.Println("asdfasdfasdfasdfasdfasdf")
				}
			//}()
		}
	}
}

func (c *client) ListenNats(name string) (err error) {
	//name = "foo"
	fmt.Println("开始监听：", name)
	_, err = ncConn.Subscribe(name, func(m *nats.Msg) {
		// 在这阻塞就完了。。。。。
		// 查询这里阻塞后不能恢复的原因
		// nats不提供缓冲，阻塞了会造成消息丢失，最大是1048575
		// 解决方案：扩大管道缓冲长度，限制流量（频道发消息间隔，个人发送消息频率，好友数量）

		c.CountFlat++
		c.Count <- c.CountFlat
		c.Send <- m.Data
	})
	if err != nil {
		fmt.Println("nats listen error：", err.Error())
	}
	return
}

type client struct {
	Send      chan []byte
	Count     chan int
	CountFlat int
}
