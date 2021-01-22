package ws

import (
	"code271/ws_test_1/pkg/const_key"
	"code271/ws_test_1/pkg/nats_helper"
	snowflake_helper "code271/ws_test_1/pkg/snowflake"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"net/http"
)

var WsUgrader = &websocket.Upgrader{ // ws协议升级
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Client is a websocket client
type Client struct {
	ID        string          // 用户id
	Ch        string          // 当前api serviceID 用作频道
	Socket    *websocket.Conn // socket
	Send      chan []byte     // 待发送消息整合
	Count     int             // log用 消息计数
	CountChan chan []int      // 消息计数队列。配合send 成对存在 目前没用上
}

// Read socket 读。获取客户端发送的信息， TODO 迁移断言逻辑到别的方法
func (c *Client) Read() {
	defer func() {
		Center.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Center.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		once := new(Message)
		if err = json.Unmarshal(message, once); err != nil {
			fmt.Println("消息反序列化失败：", err.Error())
			continue
		} else {
			fmt.Println(string(message))
		}
		/* 处理消息逻辑 */
		switch once.Type {
		case const_key.ChM: // 频道消息 本节点消息
			fmt.Println("接受到频道消息，频道是：", c.Ch)
			content := make(map[string]interface{})
			// todo 填充发送消息
			err = nats_helper.SendMessage(c.Ch, c.ID, content)
		case const_key.PersonalM: // p2p 消息

		default:
			continue
		}
		if err != nil {
			fmt.Println("发送频道消息失败：", err.Error())
		}
	}
}

// Write socket 写 获取client send 中的内容 发送到客户端
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

// ListenNats 监听nats 只有在线时才有消息接收。无缓存。 nats-streaming 拥有持久化消息
func (c *Client) ListenNats(name string) () {
	//name = "foo"
	fmt.Println("开始监听：", name)
	_, err := nats_helper.NcConn.Subscribe(name, func(m *nats.Msg) {
		c.Count++
		fmt.Println("收到消息了：", string(m.Data), "这是第：", c.Count, " 条")
		fmt.Println("队列里有: ", len(c.Send), " 条数据")
		c.Send <- m.Data
	})
	if err != nil {
		fmt.Println("nats listen error：", err.Error())
	}
	return
}

// NewClient 客户端初始化
func NewClient(id string, conn *websocket.Conn) (c *Client) {
	c = new(Client)
	if id != "" {
		c.ID = id
	} else {
		c.ID = snowflake_helper.MakeIDStr()
	}
	c.Ch = Center.ServerID
	c.Socket = conn
	c.Send = make(chan []byte, 1000000)     // 太小了会阻塞消息接收，造成消息丢失，为什么我不知道。。。
	c.CountChan = make(chan []int, 1000000) // 太小了会阻塞消息接收，造成消息丢失，为什么我不知道。。。
	return
}

//RunWsClient 开启监听事件
func RunWsClient(c *gin.Context) (err error) {
	conn, err := WsUgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := NewClient(
		"",
		conn,
	)
	Center.Register <- client
	go client.ListenNats(client.ID)
	go client.ListenNats(client.Ch)
	go client.Read()
	go client.Write()
	return
}

//GetUserWsKey 获取用户ws验证签名，准备用jwt
func GetUserWsKey(userID int64) (key string) {

	return
}
