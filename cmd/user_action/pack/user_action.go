package pack

import (
	"dy/cmd/user_action/dal/db"
	"dy/cmd/user_action/kitex_gen/useraction"
	"time"
)

func User(m *db.User, followCount int64, followerCount int64, isFollow *bool) *useraction.User {
	if m == nil {
		return nil
	}

	return &useraction.User{
		Id:            m.UserId,
		Name:          m.UserName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}
}

func Likes(ms []*db.Like) []int64 {
	vIds := make([]int64, 0)
	if len(ms) == 0 {
		return vIds
	}
	vIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			vIdMap[m.VideoId] = struct{}{}
		}
	}
	for vId := range vIdMap {
		vIds = append(vIds, vId)
	}
	return vIds
}

func Videos(ms []*db.Video) []int64 {
	vIds := make([]int64, 0)
	if len(ms) == 0 {
		return vIds
	}
	vIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			vIdMap[m.VideoId] = struct{}{}
		}
	}
	for vId := range vIdMap {
		vIds = append(vIds, vId)
	}
	return vIds
}

func Video(m *db.Video, user *useraction.User, likeCount int64, commentCount int64, isLike bool) *useraction.Video {
	if m == nil {
		return nil
	}
	return &useraction.Video{
		Id:            m.VideoId,
		Author:        user,
		PlayUrl:       m.PlayUrl,
		CoverUrl:      m.CoverUrl,
		Title:         m.Title,
		FavoriteCount: likeCount,
		CommentCount:  commentCount,
		IsFavorite:    isLike,
	}
}

func Comment(m *db.Comment, user *useraction.User) *useraction.Comment {
	if m == nil {
		return nil
	}
	return &useraction.Comment{
		Id:         m.CommentId,
		User:       user,
		Content:    m.Content,
		CreateDate: time.Unix(m.CreateDate, 0).Format("01-02"),
	}
}

func Follows(ms []*db.Follow) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.ToUserID] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}

func Followers(ms []*db.Follow) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.FromUserId] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
