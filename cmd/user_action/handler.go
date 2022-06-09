package main

import (
	"context"
	"dy/cmd/user_action/kitex_gen/useraction"
)

// UserActionServiceImpl implements the last service interface defined in the IDL.
type UserActionServiceImpl struct{}

// GetUser implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) GetUser(ctx context.Context, req *useraction.GetUserRequest) (resp *useraction.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) PublishList(ctx context.Context, req *useraction.PublishListRequest) (resp *useraction.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) FavoriteAction(ctx context.Context, req *useraction.FavoriteActionRequest) (resp *useraction.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) FavoriteList(ctx context.Context, req *useraction.FavoriteListRequest) (resp *useraction.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) CommentAction(ctx context.Context, req *useraction.CommentActionRequest) (resp *useraction.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) CommentList(ctx context.Context, req *useraction.CommentListRequest) (resp *useraction.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) RelationAction(ctx context.Context, req *useraction.RelationActionRequest) (resp *useraction.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) RelationFollowList(ctx context.Context, req *useraction.RelationFollowListRequest) (resp *useraction.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the UserActionServiceImpl interface.
func (s *UserActionServiceImpl) RelationFollowerList(ctx context.Context, req *useraction.RelationFollowerListRequest) (resp *useraction.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
