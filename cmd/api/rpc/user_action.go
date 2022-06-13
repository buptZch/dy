package rpc

import (
	"context"
	"time"

	"dy/cmd/api/kitex_gen/useraction"
	"dy/cmd/api/kitex_gen/useraction/useractionservice"
	"dy/pkg/constants"
	"dy/pkg/errno"
	"dy/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var useractionClient useractionservice.Client

func initUserActionRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := useractionservice.NewClient(
		constants.UserActionServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	useractionClient = c
}

func GetUser(ctx context.Context, req *useraction.GetUserRequest) (*useraction.User, error) {
	resp, err := useractionClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.User, nil
}

func PublishList(ctx context.Context, req *useraction.PublishListRequest) ([]*useraction.Video, error) {
	resp, err := useractionClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

func FavoriteAction(ctx context.Context, req *useraction.FavoriteActionRequest) error {
	resp, err := useractionClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func FavoriteList(ctx context.Context, req *useraction.FavoriteListRequest) ([]*useraction.Video, error) {
	resp, err := useractionClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

func CommentAction(ctx context.Context, req *useraction.CommentActionRequest) (*useraction.Comment, error) {
	resp, err := useractionClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Comment, nil
}

func CommentList(ctx context.Context, req *useraction.CommentListRequest) ([]*useraction.Comment, error) {
	resp, err := useractionClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.CommentList, nil
}

func RelationAction(ctx context.Context, req *useraction.RelationActionRequest) error {
	resp, err := useractionClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func RelationFollowList(ctx context.Context, req *useraction.RelationFollowListRequest) ([]*useraction.User, error) {
	resp, err := useractionClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserList, nil
}

func RelationFollowerList(ctx context.Context, req *useraction.RelationFollowerListRequest) ([]*useraction.User, error) {
	resp, err := useractionClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserList, nil
}
