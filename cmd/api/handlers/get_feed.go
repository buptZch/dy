package handlers

import (
	"context"
	"dy/cmd/api/kitex_gen/video"
	"dy/cmd/api/middleware"
	"dy/cmd/api/rpc"
	"dy/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"time"
)

func GetFeed(c *gin.Context) {

	var requestVar GetFeedRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendFeedResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.LatestTime == 0 {
		requestVar.LatestTime = time.Now().Unix()
	}
	//先从token中读取user_id, 为0则没有token
	uId, err := jwt.VerifyToken(requestVar.Token)
	//if err != nil {
	//	SendFeedResponse(c, nil, errno.ConvertErr(err))
	//	return
	//}
	//if uId <= 0 {
	//	klog.Error(uId)
	//	klog.Error(requestVar.Token)
	//	SendFeedResponse(c, nil, errno.ParamErr)
	//	return
	//}

	req := &video.GetFeedRequest{LatestTime: requestVar.LatestTime, UserId: uId}
	vList, err := rpc.GetFeed(context.Background(), req)
	if err != nil {
		SendFeedResponse(c, nil, errno.ConvertErr(err))
		return
	}
	klog.Errorf(string(requestVar.LatestTime), uId)
	SendFeedResponse(c, vList, nil)
	return
}
