package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
	"dy/cmd/user_action/pack"
	"time"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *CommentActionService) CommentAction(req *useraction.CommentActionRequest) (*useraction.Comment, error) {
	commentModel := &db.Comment{
		UserId:     req.UserId,
		VideoId:    req.VideoId,
		Content:    req.CommentText,
		CommentId:  req.CommentId,
		CreateDate: time.Now().Unix(),
	}

	commentModels := make([]*db.Comment, 0)
	commentModels = append(commentModels, commentModel)
	var err error
	if req.ActionType == 1 {
		err = db.CreateComment(s.ctx, commentModels)
	} else {
		err = db.DeleteComment(s.ctx, req.CommentId)
	}
	if err != nil {
		return nil, err
	}
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
	user := pack.User(modelUsers[0], followCount, followerCount, &flag)
	comment := pack.Comment(commentModel, user)
	return comment, nil
}
