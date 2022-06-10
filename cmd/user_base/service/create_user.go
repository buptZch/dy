package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"dy/pkg/errno"

	"dy/cmd/user_base/kitex_gen/userbase"

	"dy/cmd/user_base/dal/db"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *userbase.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
		UserId:   req.UserId,
	}})
}
