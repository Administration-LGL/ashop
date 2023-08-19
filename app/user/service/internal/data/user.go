package data

import (
	"context"

	v1 "ashop/api/user/service/v1"
	"ashop/app/user/service/internal/biz"
	"ashop/app/user/service/internal/data/ent"
	"ashop/app/user/service/internal/data/ent/predicate"
	"ashop/app/user/service/internal/data/ent/user"
	"ashop/app/user/service/internal/pkg/util"
	"ashop/pkg/util/snowflakegen"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, bizuser *biz.User) (*biz.User, error) {
	pwdhs, err := util.PasswordToHash(bizuser.Password)
	if err != nil {
		return nil, err
	}
	res, err := ur.data.db.User.Create().
		SetEmail(bizuser.Email).
		SetUsername(bizuser.Username).
		SetPhone(bizuser.Phone).
		SetID(snowflakegen.GenID().Int64()).
		SetPasswordHash(pwdhs).Save(ctx)
	if err != nil {
		return nil, err
	}
	bizuser.ID = res.ID
	return bizuser, nil
}

func (ur *userRepo) GetUser(ctx context.Context, bizuser *biz.User) (*biz.User, error) {
	userQuery := ur.data.db.User.Query()
	if bizuser.ID != 0 {
		userQuery = userQuery.Where(
			user.IDEQ(bizuser.ID),
		)
	}
	if len(bizuser.Email) != 0 {
		userQuery = userQuery.Where(
			user.EmailEQ(bizuser.Email),
		)
	}
	if len(bizuser.Phone) != 0 {
		userQuery = userQuery.Where(
			user.PhoneEQ(bizuser.Phone),
		)
	}
	if len(bizuser.Username) != 0 {
		userQuery = userQuery.Where(
			user.UsernameEQ(bizuser.Username),
		)
	}
	// res, err := userQuery.Where(
	// 	user.Or(user.StatusEQ(int8(v1.UserStatus_NORMAL)), user.StatusEQ(int8(v1.UserStatus_FREEZE))),
	// ).Only(ctx)
	res, err := ur.setUserQueryWithDefault_StatusOr(userQuery).Only(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("time:", res.CreatedAt)
	bizuser.ID = res.ID
	bizuser.Email = res.Email
	bizuser.Phone = res.Phone
	bizuser.Status = v1.UserStatus(res.Status)
	bizuser.Username = res.Username
	return bizuser, nil
}

func (ur *userRepo) setUserQueryWithDefault_StatusOr(userQuery *ent.UserQuery, status ...v1.UserStatus) *ent.UserQuery {
	statusLen := len(status)
	if statusLen == 0 {
		// 默认 normal/freeze
		status = make([]v1.UserStatus, 2)
		status[0] = v1.UserStatus_NORMAL
		status[1] = v1.UserStatus_FREEZE
		statusLen = len(status)
	}
	// spew.Dump(status)
	if statusLen == 1 {
		userQuery = userQuery.Where(user.StatusEQ(int8(status[0])))
	} else {
		predicateUser := make([]predicate.User, statusLen)
		for i, v := range status {
			predicateUser[i] = user.StatusEQ(int8(v))
		}
		// spew.Dump(predicateUser)
		userQuery = userQuery.Where(
			user.Or(predicateUser...),
		)
	}
	return userQuery
}

func (ur *userRepo) GetUserForAuth(ctx context.Context, phone string) (*biz.UserForAuth, error) {
	userQuery := ur.data.db.User.Query().Where(user.Phone(phone))
	userQuery = ur.setUserQueryWithDefault_StatusOr(userQuery)
	user, err := userQuery.Select(
		user.FieldID,
		user.FieldPasswordHash,
		user.FieldPhone,
		user.FieldUsername,
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.UserForAuth{
		ID:           user.ID,
		Phone:        user.Phone,
		PasswordHash: user.PasswordHash,
		Status:       v1.UserStatus(user.Status),
	}, nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, id int64, bizUser *biz.User) (*biz.User, error) {
	uuo := ur.data.db.User.UpdateOneID(id)
	if len(bizUser.Email) != 0 {
		uuo = uuo.SetEmail(bizUser.Email)
	}
	if len(bizUser.Username) != 0 {
		uuo = uuo.SetUsername(bizUser.Username)
	}
	if len(bizUser.Password) != 0 {
		pwdhs, err := util.PasswordToHash(bizUser.Password)
		if err != nil {
			return nil, err
		}
		uuo = uuo.SetPasswordHash(pwdhs)
	}
	if bizUser.Status != v1.UserStatus_UNKNOWN {
		uuo = uuo.SetStatus(int8(bizUser.Status))
	}
	user, err := uuo.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       user.ID,
		Email:    user.Email,
		Phone:    user.Phone,
		Username: user.Username,
		Status:   v1.UserStatus(user.Status),
	}, nil
}
