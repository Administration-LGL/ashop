package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

func (us *UserService) Healthy(ctx context.Context, req *empty.Empty) (reply *empty.Empty, err error) {
	return &empty.Empty{}, nil
}
