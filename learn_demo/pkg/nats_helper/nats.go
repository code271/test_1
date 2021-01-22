package nats_helper

import (
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
