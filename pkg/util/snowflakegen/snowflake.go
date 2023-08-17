package snowflakegen

import "github.com/bwmarrin/snowflake"

var globalNode *snowflake.Node

func SnowflakeNodeInit() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	globalNode = node
}

func init() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	globalNode = node
}

func GenID() snowflake.ID {
	return globalNode.Generate()
}
