package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
	"dy/cmd/user_action/pack"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *FollowListService) FollowList(req *useraction.RelationFollowListRequest) ([]*useraction.User, error) {
	modleFollows, err := db.GetFollowList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	uList := make([]*useraction.User, 0)
	UIds := pack.Follows(modleFollows)
	for _, uId := range UIds {
		uIds := make([]int64, 0)
		uIds = append(uIds, uId)
		modelUsers, err := db.MGetUser(s.ctx, uIds)
		if err != nil {
			return nil, err
		}
		followCount, err := db.GetFollowCountByUserId(s.ctx, uId)
		if err != nil {
			return nil, err
		}
		followerCount, err := db.GetFollowerCountByUserId(s.ctx, uId)
		if err != nil {
			return nil, err
		}
		flag_follow, err := db.IsFollow(s.ctx, req.UserId, uId)
		if err != nil {
			return nil, err
		}
		user := pack.User(modelUsers[0], followCount, followerCount, &flag_follow)
		uList = append(uList, user)
	}
	return uList, err
}
