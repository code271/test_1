package snowflake_helper

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init() (err error) {
	node, err = snowflake.NewNode(1)
	if err != nil {
		return
	}

	return
}

func MakeID() (id int64) {
	id = node.Generate().Int64()
	return
}

func MakeIDStr() (idStr string) {
	idStr = fmt.Sprintf("%d", node.Generate().Int64())
	return
}
