package handlers

import (
	"dy/cmd/api/kitex_gen/video"
	"mime/multipart"
	"net/http"

	"dy/pkg/errno"

	"dy/cmd/api/kitex_gen/useraction"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

type UserBaseResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
}

// SendResponse pack response for login/register
func SendUserBaseResponse(c *gin.Context, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserBaseResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserId:  userId,
		Token:   token,
	})
}

type UserParam struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

type GetUserInfoRequest struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type UserInfoResponse struct {
	Code    int64            `json:"status_code"`
	Message string           `json:"status_msg"`
	User    *useraction.User `json:"user"`
}

func SendUserInfoResponse(c *gin.Context, user *useraction.User, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		User:    user,
	})
}

type GetPublishVideoRequest struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type PublishVideoResponse struct {
	Code      int64               `json:"status_code"`
	Message   string              `json:"status_msg"`
	VideoList []*useraction.Video `json:"video_list"`
}

func SendPublishVideoResponse(c *gin.Context, vList []*useraction.Video, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, PublishVideoResponse{
		Code:      Err.ErrCode,
		Message:   Err.ErrMsg,
		VideoList: vList,
	})
}

type FavoriteActionRequest struct {
	VideoId    int64  `json:"video_id" form:"video_id"`
	Token      string `json:"token" form:"token"`
	ActionType int32  `json:"action_type" form:"action_type"`
}

type FavoriteActionResponse struct {
	Code    int64  `json:"status_status_code"`
	Message string `json:"status_msg"`
}

func SendFavoriteActionResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FavoriteActionResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

type GetFavoriteListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type FavoriteListResponse struct {
	Code      int64               `json:"status_code"`
	Message   string              `json:"status_msg"`
	VideoList []*useraction.Video `json:"video_list"`
}

func SendFavoriteListResponse(c *gin.Context, vList []*useraction.Video, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FavoriteListResponse{
		Code:      Err.ErrCode,
		Message:   Err.ErrMsg,
		VideoList: vList,
	})
}

type CommentActionRequest struct {
	VideoId     int64  `json:"video_id" form:"video_id"`
	Token       string `json:"token" form:"token"`
	ActionType  int32  `json:"action_type" form:"action_type"`
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentId   int64  `json:"comment_id" form:"comment_id"`
}

type CommentActionResponse struct {
	Code    int64               `json:"status_code"`
	Message string              `json:"status_msg"`
	Comment *useraction.Comment `json:"comment"`
}

func SendCommentActionResponse(c *gin.Context, comment *useraction.Comment, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, CommentActionResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Comment: comment,
	})
}

type GetCommentListRequest struct {
	VideoId int64  `json:"video_id" form:"video_id"`
	Token   string `json:"token" form:"token"`
}

type CommentListResponse struct {
	Code        int64                 `json:"status_code"`
	Message     string                `json:"status_msg"`
	CommentList []*useraction.Comment `json:"comment_list"`
}

func SendCommentListResponse(c *gin.Context, cList []*useraction.Comment, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, CommentListResponse{
		Code:        Err.ErrCode,
		Message:     Err.ErrMsg,
		CommentList: cList,
	})
}

type RelationActionRequest struct {
	ToUserId   int64  `json:"to_user_id" form:"to_user_id"`
	Token      string `json:"token" form:"token"`
	ActionType int32  `json:"action_type" form:"action_type"`
}

type RelationActionResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

func SendRelationActionResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RelationActionResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

type GetFollowListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type FollowListResponse struct {
	Code     int64              `json:"status_code"`
	Message  string             `json:"status_msg"`
	UserList []*useraction.User `json:"user_list"`
}

func SendFollowListResponse(c *gin.Context, uList []*useraction.User, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FollowListResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserList: uList,
	})
}

type GetFollowerListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type FollowerListResponse struct {
	Code     int64              `json:"status_code"`
	Message  string             `json:"status_msg"`
	UserList []*useraction.User `json:"user_list"`
}

func SendFollowerListResponse(c *gin.Context, uList []*useraction.User, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FollowerListResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserList: uList,
	})
}

type GetFeedRequest struct {
	LatestTime int64  `json:"latest_time" form:"latest_time"`
	Token      string `json:"token" form:"token"`
}

type FeedResponse struct {
	Code      int64          `json:"status_code"`
	Message   string         `json:"status_msg"`
	VideoList []*video.Video `json:"video_list"`
}

func SendFeedResponse(c *gin.Context, vList []*video.Video, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FeedResponse{
		Code:      Err.ErrCode,
		Message:   Err.ErrMsg,
		VideoList: vList,
	})
}

type PublishActionRequest struct {
	data  *multipart.FileHeader `json:"data" form:"data"`
	Token string                `json:"token" form:"token"`
	Title string                `json:"title" form:"title"`
}

type PublishActionResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

func SendPublishActionResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, PublishActionResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}
