package main

import (
	"context"
	"dy/cmd/user_action/kitex_gen/useraction"
	"dy/cmd/user_action/pack"
	"dy/cmd/user_action/service"
	"dy/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// UserActionServiceImpl implements the last service interface defined in the IDL.
type UserActionServiceImpl struct{}

// GetUser implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) GetUser(ctx context.Context, req *useraction.GetUserRequest) (resp *useraction.GetUserResponse, err error) {
	resp = new(useraction.GetUserResponse)

	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewGetUserInfoService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = user
	return resp, nil
}

// PublishList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) PublishList(ctx context.Context, req *useraction.PublishListRequest) (resp *useraction.PublishListResponse, err error) {
	resp = new(useraction.PublishListResponse)

	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	vList, err := service.NewGetPublishVideoService(ctx).GetPublishVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = vList
	return resp, nil
}

// FavoriteAction implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) FavoriteAction(ctx context.Context, req *useraction.FavoriteActionRequest) (resp *useraction.FavoriteActionResponse, err error) {
	resp = new(useraction.FavoriteActionResponse)

	if req.UserId <= 0 || req.VideoId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) FavoriteList(ctx context.Context, req *useraction.FavoriteListRequest) (resp *useraction.FavoriteListResponse, err error) {
	resp = new(useraction.FavoriteListResponse)
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	lList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = lList
	return resp, nil
}

// CommentAction implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) CommentAction(ctx context.Context, req *useraction.CommentActionRequest) (resp *useraction.CommentActionResponse, err error) {
	resp = new(useraction.CommentActionResponse)

	if req.UserId <= 0 || req.VideoId <= 0 { //|| req.CommentId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.ActionType == 1 && req.CommentId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	comment, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// CommentList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) CommentList(ctx context.Context, req *useraction.CommentListRequest) (resp *useraction.CommentListResponse, err error) {
	resp = new(useraction.CommentListResponse)
	if req.VideoId <= 0 {
		klog.Errorf("videoId:", req.VideoId)
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	cList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentList = cList
	return resp, nil
}

// RelationAction implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) RelationAction(ctx context.Context, req *useraction.RelationActionRequest) (resp *useraction.RelationActionResponse, err error) {
	resp = new(useraction.RelationActionResponse)

	if req.UserId <= 0 || req.ToUserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// RelationFollowList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) RelationFollowList(ctx context.Context, req *useraction.RelationFollowListRequest) (resp *useraction.RelationFollowListResponse, err error) {
	resp = new(useraction.RelationFollowListResponse)
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	uList, err := service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = uList
	return resp, nil
}

// RelationFollowerList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) RelationFollowerList(ctx context.Context, req *useraction.RelationFollowerListRequest) (resp *useraction.RelationFollowerListResponse, err error) {
	resp = new(useraction.RelationFollowerListResponse)
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	uList, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = uList
	return resp, nil
}
