// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"bytes"
	"context"
	"dy/cmd/video/kitex_gen/video"
	"dy/pkg/constants"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"os"
	"time"
)

type Video struct {
	VideoId    int64  `json:"video_id"`
	UserId     int64  `json:"user_id"`
	PlayUrl    string `json:"play_url"`
	Title      string `json:"title"`
	CoverUrl   string `json:"cover_url"`
	CreateDate int64  `json:"create_date"`
}
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

const srcpath = "../../src"

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func (n *Comment) TableName() string {
	return constants.CommentTableName
}

func (n *Like) TableName() string {
	return constants.LikeTableName
}

func (n *Follow) TableName() string {
	return constants.FollowTableName
}

// QueryUser query list of user info
func TopVideo(ctx context.Context, LastTime int64, userId int64) ([]*video.Video, int64, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("create_date < ?", LastTime).Limit(30).Order("create_date desc").Find(&res).Error; err != nil {
		return nil, time.Now().Unix(), err
	}
	result := make([]*video.Video, 0)
	for _, v := range res {
		fcount, err := GetLikeCountByVideoId(ctx, v.VideoId)
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		ccount, err := GetCommentCountByVideoId(ctx, v.VideoId)
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		ifF, err := IsLikeByUserId(ctx, userId, v.VideoId)
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		vuName, err := MGetUser(ctx, v.UserId)
		klog.Errorf("vUserId", vuName+" "+string(v.UserId))
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		vufollowc, err := GetFollowCountByUserId(ctx, v.UserId)
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		vufollowerc, err := GetFollowerCountByUserId(ctx, v.UserId)
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		isFollowByU, err := IsFollow(ctx, userId, v.UserId)
		if err != nil {
			return nil, time.Now().Unix(), err
		}
		videoUser := video.User{
			Id:            v.UserId,
			Name:          vuName,
			FollowCount:   vufollowc,
			FollowerCount: vufollowerc,
			IsFollow:      &isFollowByU,
		}
		tempVideo := video.Video{
			Id:            v.VideoId,
			Author:        &videoUser,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: fcount,
			CommentCount:  ccount,
			IsFavorite:    ifF,
		}
		result = append(result, &tempVideo)
	}
	next_time := time.Now().Unix()
	if len(res) > 0 {
		next_time = res[0].CreateDate
	}
	return result, next_time, nil
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

//是否已经关注
func IsFollow(ctx context.Context, fromUserId int64, toUserId int64) (bool, error) {
	res := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("from_user_id = ? and to_user_id = ? ", fromUserId, toUserId).Find(&res).Error; err != nil {
		return false, err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

//
//func IsFollow(ctx context.Context, fromUserId int64, toUserId int64) (bool, error) {
//	res := make([]*Follow, 0)
//	if err := DB.WithContext(ctx).Where("from_user_id = ? and to_user_id = ? ", fromUserId, toUserId).Find(&res).Error; err != nil {
//		return false, err
//	}
//	klog.Errorf("关注：", res)
//	if len(res) == 0 {
//		return false, nil
//	}
//	return true, nil
//}

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

//获取用户Name
func MGetUser(ctx context.Context, userID int64) (string, error) {
	var u = User{UserId: userID}
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&u).Error; err != nil {
		return "", err
	}
	return u.UserName, nil
}

//上传视频
func UploadVideo(ctx context.Context, data []byte, id int64, videoId int64, title string) (bool, error) {

	var unixTime int64 = time.Now().Unix()
	var (
		fileName = string(id) + string(unixTime) + ".mp4"
		content  = data
		err      error
	)
	klog.Errorf("filename:", fileName)
	if err = ioutil.WriteFile(srcpath+fileName, []byte(content), 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return false, err
	}
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(fileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return false, err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return false, err
	}
	imgName := string(id) + string(unixTime) + ".jpg"
	klog.Errorf("imgUrl", srcpath+imgName)
	err = imaging.Save(img, srcpath+imgName)
	if err != nil {
		return false, err
	}
	ip := "10.112.197.219:8001/"
	videoInstance := Video{
		VideoId:    videoId,
		UserId:     id,
		PlayUrl:    fileName,
		Title:      title,
		CoverUrl:   ip + imgName,
		CreateDate: unixTime,
	}
	DB.Create(videoInstance)
	return true, nil
}
