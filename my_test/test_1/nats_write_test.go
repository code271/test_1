package test_1_test

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}

func TestSendNats(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("nats 链接错误：", err.Error())
		return
	}
	name := "foo"
	// 发布消息
	//ListenIt(nc, name)
	wg.Add(2)
	go func() {
		for i := 0; i < 5000; i++ {
			err = nc.Publish(name, []byte("Hello f##k"))
			//wg.Add(1)
			if err != nil {
				t.Log("error:", err.Error())
			}
		}
		fmt.Println(" 发送成功！")
		wg.Done()
	}()
	go func() {
		for i := 0; i < 5000; i++ {
			err = nc.Publish(name, []byte("Hello f##k"))
			//wg.Add(1)
			if err != nil {
				t.Log("error:", err.Error())
			}
		}
		fmt.Println(" 发送成功！")
		wg.Done()
	}()
	wg.Wait()
	defer nc.Close()
	return
}

//func ListenIt(nc *nats.Conn, subName string) {
//	_, err := nc.Subscribe(subName, func(m *nats.Msg) {
//		wg.Done()
//		fmt.Println("收到消息了：", string(m.Data))
//	})
//	if err != nil {
//		fmt.Println("监听消息失败：", err.Error())
//		panic(err.Error())
//	}
//	return
//}

func TestNatsAnswer(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("nats 链接错误：", err.Error())
		return
	}
	_, _ = nc.Subscribe("help", func(m *nats.Msg) {
		err = nc.Publish(m.Reply, []byte("fuck I can help!"))
		fmt.Println("收到了什么东西: ", string(m.Data))
		if err != nil {
			t.Log("reply error :", err.Error())
			return
		}
	})

	msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		t.Log("reply error :", err.Error())
		return
	}
	fmt.Println(string(msg.Data))
	return
}

func TestNatsAnswer2(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("nats 链接错误：", err.Error())
		return
	}
	_, _ = nc.Subscribe("help", func(m *nats.Msg) {
		err = m.Respond([]byte("fuck I can help!"))
		fmt.Println("收到了什么东西: ", string(m.Data))
		if err != nil {
			t.Log("reply error :", err.Error())
			return
		}
	})

	msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		t.Log("reply error :", err.Error())
		return
	}
	fmt.Println(string(msg.Data))
	return
}

func TestGetOnlineUser(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("nats 链接错误：", err.Error())
		return
	}
	name := "server_center"
	data := new(NcRequest)
	data.Answer = "get_online"
	buf, _ := json.Marshal(data)
	msg, err := nc.Request(name, buf, 10*time.Millisecond)
	if err != nil {
		t.Log("nats 请求错误：", err.Error())
		return
	}
	fmt.Println("得到的答复：", string(msg.Data))

	if len(msg.Data) == 0 {
		fmt.Println("返回消息异常")
		return
	}

	rep := &struct {
		OnlineUser []string `json:"online_user"`
	}{}

	_ = json.Unmarshal(msg.Data, rep)
	fmt.Println("获取到的在线用户：", rep.OnlineUser)

	// 欢迎
	for _, v := range rep.OnlineUser {
		_ = nc.Publish(v, []byte("hello "+v))
		_ = nc.Publish(v, []byte("nice to meet you!"))
	}

	return
}

type NcRequest struct {
	Answer string      `json:"answer"`
	Data   interface{} `json:"data"`
}

// 多个回复怎么办 queue模式，但是无法一次性收到所有人的回复
func TestNatsAnswer3(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Log("nats 链接错误：", err.Error())
		return
	}
	_, _ = nc.QueueSubscribe("help", "test", func(m *nats.Msg) {
		err = m.Respond([]byte("fuck I can help 1 !"))
		fmt.Println("收到了什么东西1: ", string(m.Data))
		if err != nil {
			t.Log("reply error :", err.Error())
			return
		}
	})

	_, _ = nc.QueueSubscribe("help", "test", func(m *nats.Msg) {
		err = m.Respond([]byte("fuck I can help 2 !"))
		fmt.Println("收到了什么东西2: ", string(m.Data))
		if err != nil {
			t.Log("reply error :", err.Error())
			return
		}
	})

	_, _ = nc.QueueSubscribe("help", "test", func(m *nats.Msg) {
		err = m.Respond([]byte("fuck I can help 3 !"))
		fmt.Println("收到了什么东西3: ", string(m.Data))
		if err != nil {
			t.Log("reply error :", err.Error())
			return
		}
	})

	for i := 0; i < 3; i++ {
		msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
		if err != nil {
			t.Log("reply error :", err.Error())
			return
		}
		fmt.Println(string(msg.Data))
	}

	return
}

func TestGetSendMessage(t *testing.T) {
	a := new(Message)
	a.Type = "ch_m"
	a.Sender = "1234512345"
	a.Content = "这是一个频道消息"
	buf, _ := json.Marshal(a)
	fmt.Println(string(buf))
	return
}

type Message struct {
	Sender    string `json:"sender,omitempty"`    // 发送者
	Type      string `json:"type"`                // 消息类别
	Recipient string `json:"recipient,omitempty"` // 接受者
	Content   string `json:"content,omitempty"`   // 消息内容
}
