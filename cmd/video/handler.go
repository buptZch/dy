package main

import (
	"context"
	"dy/cmd/video/kitex_gen/video"
	"dy/cmd/video/pack"
	"dy/cmd/video/service"
	"dy/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.GetFeedRequest) (resp *video.GetFeedResponse, err error) {
	resp = new(video.GetFeedResponse)

	res, nextTime, err := service.NewFeedService(ctx).GetFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = res
	resp.NextTime = nextTime
	return resp, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	resp = new(video.PublishActionResponse)

	status, err := service.NewActionService(ctx).PublishAction(req)
	if !status {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
