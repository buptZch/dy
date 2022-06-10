// Code generated by Kitex v0.3.2. DO NOT EDIT.

package useractionservice

import (
	"context"
	"dy/cmd/api/kitex_gen/useraction"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUser(ctx context.Context, req *useraction.GetUserRequest, callOptions ...callopt.Option) (r *useraction.GetUserResponse, err error)
	PublishList(ctx context.Context, req *useraction.PublishListRequest, callOptions ...callopt.Option) (r *useraction.PublishListResponse, err error)
	FavoriteAction(ctx context.Context, req *useraction.FavoriteActionRequest, callOptions ...callopt.Option) (r *useraction.FavoriteActionResponse, err error)
	FavoriteList(ctx context.Context, req *useraction.FavoriteListRequest, callOptions ...callopt.Option) (r *useraction.FavoriteListResponse, err error)
	CommentAction(ctx context.Context, req *useraction.CommentActionRequest, callOptions ...callopt.Option) (r *useraction.CommentActionResponse, err error)
	CommentList(ctx context.Context, req *useraction.CommentListRequest, callOptions ...callopt.Option) (r *useraction.CommentListResponse, err error)
	RelationAction(ctx context.Context, req *useraction.RelationActionRequest, callOptions ...callopt.Option) (r *useraction.RelationActionResponse, err error)
	RelationFollowList(ctx context.Context, req *useraction.RelationFollowListRequest, callOptions ...callopt.Option) (r *useraction.RelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, req *useraction.RelationFollowerListRequest, callOptions ...callopt.Option) (r *useraction.RelationFollowerListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserActionServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserActionServiceClient struct {
	*kClient
}

func (p *kUserActionServiceClient) GetUser(ctx context.Context, req *useraction.GetUserRequest, callOptions ...callopt.Option) (r *useraction.GetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, req)
}

func (p *kUserActionServiceClient) PublishList(ctx context.Context, req *useraction.PublishListRequest, callOptions ...callopt.Option) (r *useraction.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, req)
}

func (p *kUserActionServiceClient) FavoriteAction(ctx context.Context, req *useraction.FavoriteActionRequest, callOptions ...callopt.Option) (r *useraction.FavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, req)
}

func (p *kUserActionServiceClient) FavoriteList(ctx context.Context, req *useraction.FavoriteListRequest, callOptions ...callopt.Option) (r *useraction.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, req)
}

func (p *kUserActionServiceClient) CommentAction(ctx context.Context, req *useraction.CommentActionRequest, callOptions ...callopt.Option) (r *useraction.CommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, req)
}

func (p *kUserActionServiceClient) CommentList(ctx context.Context, req *useraction.CommentListRequest, callOptions ...callopt.Option) (r *useraction.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, req)
}

func (p *kUserActionServiceClient) RelationAction(ctx context.Context, req *useraction.RelationActionRequest, callOptions ...callopt.Option) (r *useraction.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}

func (p *kUserActionServiceClient) RelationFollowList(ctx context.Context, req *useraction.RelationFollowListRequest, callOptions ...callopt.Option) (r *useraction.RelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowList(ctx, req)
}

func (p *kUserActionServiceClient) RelationFollowerList(ctx context.Context, req *useraction.RelationFollowerListRequest, callOptions ...callopt.Option) (r *useraction.RelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowerList(ctx, req)
}
