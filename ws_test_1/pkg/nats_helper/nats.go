package nats_helper

import (
	"code271/ws_test_1/pkg/const_key"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"sync"
)

var natsOnce sync.Once
var NcConn *nats.Conn

func NatsInit() (err error) {
	natsOnce.Do(func() {
		ncFirst, errFirst := nats.Connect(nats.DefaultURL)
		if errFirst != nil {
			err = errFirst
			return
		}
		NcConn = ncFirst
	})
	return
}

func SendMessage(recipient, sender string, content map[string]interface{}) (err error) {
	data := make(map[string]interface{})
	data["sender"] = sender
	data["recipient"] = recipient
	data["content"] = content
	buf, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = NcConn.Publish(recipient, buf)
	return
}

func SendMessageToUser(user, message, sender string) (err error) {
	data := &struct {
		Sender    string `json:"sender,omitempty"`    // 发送者
		Type      string `json:"type"`                // 消息类别
		Recipient string `json:"recipient,omitempty"` // 接受者
		Content   string `json:"content,omitempty"`   // 消息内容
	}{}
	data.Type = const_key.PersonalM
	return
}
