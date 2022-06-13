package db

import (
	"context"
	"dy/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Comment struct {
	//gorm.Model
	UserId     int64  `json:"user_id"`
	CommentId  int64  `json:"comment_id"`
	Content    string `json:"content"`
	CreateDate int64  `json:"create_date"`
	VideoId    int64  `json:"video_id"`
}

type User struct {
	//gorm.Model
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Video struct {
	VideoId    int64  `json:"video_id"`
	UserID     int64  `json:"user_id"`
	PlayUrl    string `json:"play_url"`
	Title      string `json:"title"`
	CoverUrl   string `json:"cover_url"`
	CreateDate int64  `json:"create_date"`
}

type Like struct {
	Key     int32 `json:"key"`
	VideoId int64 `json:"video_id"`
	UserID  int64 `json:"user_id"`
}

type Follow struct {
	Key        int32 `json:"key"`
	FromUserId int64 `json:"from_user_id"`
	ToUserID   int64 `json:"to_user_id"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func (n *Comment) TableName() string {
	return constants.CommentTableName
}

func (n *Video) TableName() string {
	return constants.VideoTableName
}

func (n *Like) TableName() string {
	return constants.LikeTableName
}

func (n *Follow) TableName() string {
	return constants.FollowTableName
}

//获取用户info
func MGetUser(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)

	if err := DB.WithContext(ctx).Where("user_id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//赞
func LikeVideo(ctx context.Context, likes []*Like) error {
	if err := DB.WithContext(ctx).Create(likes).Error; err != nil {
		return err
	}
	return nil
}

//取消赞
func UnlikeVideo(ctx context.Context, userId int64, videoId int64) error {
	return DB.WithContext(ctx).Where("video_id = ? and user_id = ? ", videoId, userId).Delete(&Like{}).Error
}

func GetPublishVideoId(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//通过当前用户的id查询所有赞的videoId
func GetLikeVideoId(ctx context.Context, userId int64) ([]*Like, error) {
	res := make([]*Like, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//批量通过视频videoId获取Video
func MGetVideoByVideoId(ctx context.Context, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("video_id in ?", videoIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//一些端口如何进行批量查询

//查询登录用户对视频是否点赞
func IsLikeByUserId(ctx context.Context, userId int64, videoId int64) (bool, error) {
	res := make([]*Like, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ? ", userId, videoId).Find(&res).Error; err != nil {
		return false, err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

//查询视频的点赞总数
func GetLikeCountByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var total int64
	conn := DB.WithContext(ctx).Model(&Like{}).Where("video_id = ?", videoId)

	if err := conn.Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

//查询视频的评论总数
func GetCommentCountByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var total int64
	conn := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", videoId)

	if err := conn.Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

//是否已经关注
func IsFollow(ctx context.Context, fromUserId int64, toUserId int64) (bool, error) {
	res := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("from_user_id = ? and to_user_id = ? ", fromUserId, toUserId).Find(&res).Error; err != nil {
		return false, err
	}
	klog.Errorf("关注：", res)
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

//获得用户的粉丝总数
func GetFollowerCountByUserId(ctx context.Context, ToUserId int64) (int64, error) {
	var total int64
	conn := DB.WithContext(ctx).Model(&Follow{}).Where("to_user_id = ?", ToUserId)

	if err := conn.Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

//获得用户的关注总数
func GetFollowCountByUserId(ctx context.Context, FromUserId int64) (int64, error) {
	var total int64
	conn := DB.WithContext(ctx).Model(&Follow{}).Where("from_user_id = ?", FromUserId)

	if err := conn.Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

//创建评论
func CreateComment(ctx context.Context, comments []*Comment) error {
	if err := DB.WithContext(ctx).Create(comments).Error; err != nil {
		return err
	}
	return nil
}

//删除评论
func DeleteComment(ctx context.Context, commentId int64) error {
	return DB.WithContext(ctx).Where("comment_id = ? ", commentId).Delete(&Comment{}).Error
}

//通过视频videoId获取CommentList
func GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//关注
func FollowUser(ctx context.Context, follows []*Follow) error {
	if err := DB.WithContext(ctx).Create(follows).Error; err != nil {
		return err
	}
	return nil
}

//取消关注
func UnFollowUser(ctx context.Context, fromUserId int64, toUserId int64) error {
	return DB.WithContext(ctx).Where("from_user_id = ? and to_user_id = ? ", fromUserId, toUserId).Delete(&Follow{}).Error
}

//关注列表
func GetFollowList(ctx context.Context, userId int64) ([]*Follow, error) {
	res := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("from_user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//粉丝列表
func GetFollowerList(ctx context.Context, userId int64) ([]*Follow, error) {
	res := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
