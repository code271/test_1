package ws

import (
	"code271/learn_demo/pkg/nats_helper"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"

)

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		fmt.Println("收到了什么东西：", string(jsonMessage))
		Manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) ListenNats(name string) () {
	_, err := nats_helper.NcConn.Subscribe(name, func(m *nats.Msg) {
		fmt.Println("收到消息了：", string(m.Data))
		c.Send <- m.Data
		fmt.Println("已经扔进队列了！")
	})
	if err != nil {
		fmt.Println("nats listen error：", err.Error())
	}
	return
}

func (c *Client) NewClient(id string, conn *websocket.Conn) () {
	c.ID = id
	c.Socket = conn
	c.Send = make(chan []byte, 1000000)
	return
}
