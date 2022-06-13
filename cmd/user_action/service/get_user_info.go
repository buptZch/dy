package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
	"dy/cmd/user_action/pack"
)

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *GetUserInfoService) GetUser(req *useraction.GetUserRequest) (*useraction.User, error) {
	uIds := make([]int64, 0)
	uIds = append(uIds, req.UserId)
	modelUsers, err := db.MGetUser(s.ctx, uIds)
	if err != nil {
		return nil, err
	}
	followCount, err := db.GetFollowCountByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	followerCount, err := db.GetFollowerCountByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	flag, err := db.IsFollow(s.ctx, req.UserId, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.User(modelUsers[0], followCount, followerCount, &flag), nil
}
