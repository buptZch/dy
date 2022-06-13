package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
	"dy/cmd/user_action/pack"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *CommentListService) CommentList(req *useraction.CommentListRequest) ([]*useraction.Comment, error) {

	modelComments, err := db.GetCommentListByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	cList := make([]*useraction.Comment, 0)
	for _, modelComment := range modelComments {
		uId := modelComment.UserId
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

		flag, err := db.IsFollow(s.ctx, req.UserId, uId)

		if err != nil {
			return nil, err
		}
		user := pack.User(modelUsers[0], followCount, followerCount, &flag)
		comment := pack.Comment(modelComment, user)
		cList = append(cList, comment)
	}
	return cList, err
}
