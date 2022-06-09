package main

import (
	"context"
	"dy/cmd/user_base/kitex_gen/userbase"
	"dy/cmd/user_base/pack"
	"dy/cmd/user_base/service"
	"dy/pkg/errno"
)

// UserBaseServiceImpl implements the last service interface defined in the IDL.
type UserBaseServiceImpl struct{}

// CreateUser implements the UserBaseServiceImpl interface.
func (s *UserBaseServiceImpl) CreateUser(ctx context.Context, req *userbase.CreateUserRequest) (resp *userbase.CreateUserResponse, err error) {
	resp = new(userbase.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserBaseServiceImpl interface.
func (s *UserBaseServiceImpl) CheckUser(ctx context.Context, req *userbase.CheckUserRequest) (resp *userbase.CheckUserResponse, err error) {
	resp = new(userbase.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
