package handlers

import (
	"context"
	"dy/cmd/api/kitex_gen/useraction"
	"dy/cmd/api/middleware"
	"dy/cmd/api/rpc"
	"dy/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func GetPublishVideo(c *gin.Context) {

	var requestVar GetPublishVideoRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendPublishVideoResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	_, err := jwt.VerifyToken(requestVar.Token)
	if err != nil {
		SendPublishVideoResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.UserId <= 0 {
		klog.Error(requestVar.UserId)
		klog.Error(requestVar.Token)
		SendPublishVideoResponse(c, nil, errno.ParamErr)
		return
	}

	req := &useraction.PublishListRequest{UserId: requestVar.UserId}
	vList, err := rpc.PublishList(context.Background(), req)
	if err != nil {
		SendPublishVideoResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendPublishVideoResponse(c, vList, nil)
	return
}
