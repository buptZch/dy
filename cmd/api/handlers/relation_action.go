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

func RelationAction(c *gin.Context) {

	var requestVar RelationActionRequest

	if err := c.ShouldBind(&requestVar); err != nil {
		SendRelationActionResponse(c, errno.ConvertErr(err))
		return
	}
	//先验证token
	uId, err := jwt.VerifyToken(requestVar.Token)
	klog.Errorf("uid", uId)
	if err != nil {
		SendRelationActionResponse(c, errno.ConvertErr(err))
		return
	}
	if uId <= 0 {
		klog.Error(uId)
		klog.Error(requestVar.Token)
		SendRelationActionResponse(c, errno.ParamErr)
		return
	}

	req := &useraction.RelationActionRequest{UserId: uId, ToUserId: requestVar.ToUserId, ActionType: requestVar.ActionType}
	err = rpc.RelationAction(context.Background(), req)
	if err != nil {
		SendRelationActionResponse(c, errno.ConvertErr(err))
		return
	}
	SendRelationActionResponse(c, nil)
	return
}
