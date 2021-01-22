package ws


// Message is an object for websocket message which is mapped to json type
type Message struct {
	Type      string `json:"type"`                // 消息类别
	Recipient string `json:"recipient,omitempty"` // 接受者
	Content   string `json:"content,omitempty"`   // 消息内容
}
