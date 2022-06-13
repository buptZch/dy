package service

import (
	"context"
	"dy/cmd/video/dal/db"
	"dy/cmd/video/kitex_gen/video"
)

type FeedService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

// Video Service
func (s *FeedService) GetFeed(req *video.GetFeedRequest) ([]*video.Video, int64, error) {
	res, nextTime, err := db.TopVideo(s.ctx, req.LatestTime, req.UserId)
	if err != nil {
		return nil, nextTime, err
	}
	return res, nextTime, nil
}
