package biz

import (
	"context"
	"net/http"
	"strings"

	v1 "ashop/api/user/service/v1"
	"ashop/app/user/service/internal/data/ent"
	validphonenum "ashop/pkg/util/valid_phone_num"

	"ashop/app/user/service/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.UserServiceErrorReason_NOT_FOUND_ERROR.String(), "user not found")
)

// Greeter is a Greeter model.
type User struct {
	ID       uint64
	Username string
	Phone    string
	Email    string
	Password string
	Status   v1.UserStatus
}
type UserForAuth struct {
	ID           uint64
	Phone        string
	PasswordHash string
	Status       v1.UserStatus
}

// UserRepo is a Greater repo.
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, user *User) (*User, error)
	GetUserForAuth(ctx context.Context, phone string) (*UserForAuth, error)
}

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) GetUser(ctx context.Context, req *v1.GetUserReq) (*User, error) {
	user, err := uc.ur.GetUser(ctx, &User{
		ID:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, errors.New(http.StatusNotAcceptable, "getuser", err.Error())
	}
	return user, nil
}

func (uc *UserUsecase) checkUserRegisterReqInfo(ctx context.Context, user *User) error {
	errHead := "register baseinfo"
	if len(strings.TrimSpace(user.Username)) == 0 ||
		len(strings.TrimSpace(user.Password)) == 0 ||
		len(strings.TrimSpace(user.Phone)) == 0 {
		return errors.New(http.StatusNotAcceptable, errHead, "username/password/phone must be bot null")
	}
	// 校验手机号是否合格
	if !validphonenum.FooCheckIsNum(user.Phone) {
		return errors.New(http.StatusNotAcceptable, errHead, "not a phone num")
	}
	return nil
}

func (uc *UserUsecase) Register(ctx context.Context, user *User) (*User, error) {
	if err := uc.checkUserRegisterReqInfo(ctx, user); err != nil {
		return nil, err
	}
	// 检查是否存在Phone用户状态位正常或冻结
	_, err := uc.ur.GetUser(ctx, &User{Phone: user.Phone})
	if err == nil || !ent.IsNotFound(err) {
		return nil, errors.New(http.StatusInternalServerError, "createuser", "the phone has binding user")
	}

	user, err = uc.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "createuser", err.Error())
	}
	return user, nil
}

func (uc *UserUsecase) Login(ctx context.Context, phone, password string) (bool, error) {
	user, err := uc.ur.GetUserForAuth(ctx, phone)
	if err != nil {
		return false, errors.New(http.StatusInternalServerError, "body", err.Error())
	}
	if user.Status == v1.UserStatus_FREEZE {
		// 账号冻结
		return false, errors.New(http.StatusInternalServerError, "body", "user is freezed")
	}
	// 验证密码正确性
	return util.VertifyPasswordHash(user.PasswordHash, password), nil
}
