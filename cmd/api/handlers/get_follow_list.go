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

func GetFollowList(c *gin.Context) {

	var requestVar GetFollowListRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendFollowListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	_, err := jwt.VerifyToken(requestVar.Token)
	if err != nil {
		SendFollowListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.UserId <= 0 {
		klog.Error(requestVar.UserId)
		klog.Error(requestVar.Token)
		SendFollowListResponse(c, nil, errno.ParamErr)
		return
	}

	req := &useraction.RelationFollowListRequest{UserId: requestVar.UserId}
	uList, err := rpc.RelationFollowList(context.Background(), req)
	if err != nil {
		SendFollowListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendFollowListResponse(c, uList, nil)
	return
}
