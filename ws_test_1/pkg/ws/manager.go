package ws

import (
	"encoding/json"
	"fmt"
)

// ClientManager is a websocket manager
type ClientCenter struct {
	ServerID   string
	Clients    map[string]bool // 全局客户端在线情况
	Register   chan *Client    // 注册队列
	Unregister chan *Client    // 注销队列
}

// Manager define a ws server manager
var Center = ClientCenter{
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]bool),
}

// Start is to start a ws server
func (manager *ClientCenter) Start(serverID string) {
	Center.ServerID = serverID
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn.ID] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			conn.Send <- jsonMessage
			fmt.Println("用户：", conn.ID, " 已经登录")
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn.ID]; ok {
				close(conn.Send)
				delete(manager.Clients, conn.ID)
				fmt.Println("用户：", conn.ID, " 已经离线")
			}
		}
	}
}
