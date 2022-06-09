package main

import (
	"context"
	"dy/cmd/user_base/kitex_gen/userbase"
)

// UserBaseServiceImpl implements the last service interface defined in the IDL.
type UserBaseServiceImpl struct{}

// CreateUser implements the UserBaseServiceImpl interface.
func (s *UserBaseServiceImpl) CreateUser(ctx context.Context, req *userbase.CreateUserRequest) (resp *userbase.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserBaseServiceImpl interface.
func (s *UserBaseServiceImpl) CheckUser(ctx context.Context, req *userbase.CheckUserRequest) (resp *userbase.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}
