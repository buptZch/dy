package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *FavoriteActionService) FavoriteAction(req *useraction.FavoriteActionRequest) error {
	likeModel := &db.Like{
		UserID:  req.UserId,
		VideoId: req.VideoId,
	}
	likeModels := make([]*db.Like, 0)
	likeModels = append(likeModels, likeModel)
	var err error
	if req.ActionType == 1 {
		err = db.LikeVideo(s.ctx, likeModels)
	} else {
		err = db.UnlikeVideo(s.ctx, likeModel.UserID, likeModel.VideoId)
	}
	if err != nil {
		return err
	}
	return err
}
