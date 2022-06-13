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

func GetUser(c *gin.Context) {

	var requestVar GetUserInfoRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendUserInfoResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	_, err := jwt.VerifyToken(requestVar.Token)
	if err != nil {
		SendUserInfoResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.UserId <= 0 {
		klog.Error(requestVar.UserId)
		klog.Error(requestVar.Token)
		SendUserInfoResponse(c, nil, errno.ParamErr)
		return
	}

	req := &useraction.GetUserRequest{UserId: requestVar.UserId}
	user, err := rpc.GetUser(context.Background(), req)
	if err != nil {
		SendUserInfoResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendUserInfoResponse(c, user, nil)
	return
}
