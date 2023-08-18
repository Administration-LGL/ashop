package service

import (
	v1 "ashop/api/user/service/v1"
	"ashop/app/user/service/internal/biz"
	"context"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (us *UserService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	user, err := us.uc.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{Id: user.ID}, nil
}

func (us *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	bizUser, err := us.uc.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id:         bizUser.ID,
		Phone:      bizUser.Phone,
		Email:      bizUser.Email,
		Username:   bizUser.Username,
		UserStatus: bizUser.Status,
	}, nil
}

func (us *UserService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	result, err := us.uc.Login(ctx, req.Phone, req.Password)
	if err != nil {
		return nil, err
	}

	return &v1.LoginReply{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}
func (us *UserService) SetUserStatusByID(ctx context.Context, req *v1.SetUserStatusByIDReq) (*v1.SetUserStatusReply, error) {
	result, err := us.uc.SetUserStatus(ctx, req.Id, req.Status)
	if err != nil {
		return nil, err
	}
	return &v1.SetUserStatusReply{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}
