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

func GetFollowerList(c *gin.Context) {

	var requestVar GetFollowerListRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendFollowerListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	_, err := jwt.VerifyToken(requestVar.Token)
	if err != nil {
		SendFollowerListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.UserId <= 0 {
		klog.Error(requestVar.UserId)
		klog.Error(requestVar.Token)
		SendFollowerListResponse(c, nil, errno.ParamErr)
		return
	}

	req := &useraction.RelationFollowerListRequest{UserId: requestVar.UserId}
	uList, err := rpc.RelationFollowerList(context.Background(), req)
	if err != nil {
		SendFollowerListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendFollowerListResponse(c, uList, nil)
	return
}
