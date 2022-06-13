package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *RelationActionService) RelationAction(req *useraction.RelationActionRequest) error {
	relationModel := &db.Follow{
		FromUserId: req.UserId,
		ToUserID:   req.ToUserId,
	}
	relationModels := make([]*db.Follow, 0)
	relationModels = append(relationModels, relationModel)
	var err error
	if req.ActionType == 1 {
		err = db.FollowUser(s.ctx, relationModels)
	} else {
		err = db.UnFollowUser(s.ctx, req.UserId, req.ToUserId)
	}
	if err != nil {
		return err
	}
	return err
}
