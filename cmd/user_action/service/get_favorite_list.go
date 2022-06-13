package service

import (
	"context"
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
	"dy/cmd/user_action/pack"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

//服务里把response所有需要的返回值送入handler中装配
func (s *FavoriteListService) FavoriteList(req *useraction.FavoriteListRequest) ([]*useraction.Video, error) {

	modelLikes, err := db.GetLikeVideoId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	vList := make([]*useraction.Video, 0)
	vIds := pack.Likes(modelLikes)
	modelVideos, err := db.MGetVideoByVideoId(s.ctx, vIds)
	for _, modelVideo := range modelVideos {
		vId := modelVideo.VideoId
		uId := modelVideo.UserID
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
		flagLike, err := db.IsLikeByUserId(s.ctx, req.UserId, vId)
		if err != nil {
			return nil, err
		}
		likeCount, err := db.GetLikeCountByVideoId(s.ctx, vId)
		if err != nil {
			return nil, err
		}
		commentCount, err := db.GetCommentCountByVideoId(s.ctx, vId)
		if err != nil {
			return nil, err
		}
		video := pack.Video(modelVideo, user, likeCount, commentCount, flagLike)
		vList = append(vList, video)
	}
	return vList, err
}
