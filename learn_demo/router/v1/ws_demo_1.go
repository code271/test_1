package v1

import (
	"code271/learn_demo/pkg/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func WsPage(c *gin.Context) {
	// change the reqest to websocket model
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}

	ws.Manager.Register <- client

	 client.ListenNats("foo")

	go client.Read()
	go client.Write()
}
