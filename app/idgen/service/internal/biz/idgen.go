package biz

import (
	"ashop/app/idgen/service/internal/conf"
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// IdGenUsecase is a Greeter usecase.
type IdGenUsecase struct {
	log *log.Helper
	sn  *snowflake.Node
}

func InitSnowflake(nodeConf *conf.Node) *snowflake.Node {
	node, err := snowflake.NewNode(nodeConf.Id)
	if err != nil {
		panic(err)
	}
	return node
}

// NewIdGenUsecase new a Greeter usecase.
func NewIdGenUsecase(logger log.Logger, sn *snowflake.Node) *IdGenUsecase {
	return &IdGenUsecase{log: log.NewHelper(logger), sn: sn}
}

func (ig *IdGenUsecase) GenerateID(ctx context.Context) int64 {
	return ig.sn.Generate().Int64()
}
func (ig *IdGenUsecase) GenerateIDs(ctx context.Context, count int64) []int64 {
	ids := make([]int64, count)
	var i int64
	for i = 0; i < count; i++ {
		ids[i] = ig.sn.Generate().Int64()
	}
	return ids
}
