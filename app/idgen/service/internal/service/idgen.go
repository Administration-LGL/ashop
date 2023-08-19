package service

import (
	v1 "ashop/api/idgen/service/v1"
	"ashop/app/idgen/service/internal/biz"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

// IdGenService is a IdGen service.
type IdGenService struct {
	v1.UnimplementedIdGenServer
	ig *biz.IdGenUsecase
}

// NewIdGenService new a IdGen service.
func NewIdGenService(ig *biz.IdGenUsecase) *IdGenService {
	return &IdGenService{ig: ig}
}

func (is *IdGenService) GenerateID(ctx context.Context, req *empty.Empty) (*v1.GenIDReply, error) {
	return &v1.GenIDReply{Id: is.ig.GenerateID(ctx)}, nil
}
func (is *IdGenService) GenerateIDs(ctx context.Context, req *v1.GenIDsReq) (*v1.GenIDsReply, error) {
	return &v1.GenIDsReply{
		Ids: is.ig.GenerateIDs(ctx, req.Count),
	}, nil
}
