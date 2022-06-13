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

func GetCommentList(c *gin.Context) {

	var requestVar GetCommentListRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendCommentListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	uId, err := jwt.VerifyToken(requestVar.Token)
	if err != nil {
		SendCommentListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.VideoId <= 0 || uId < 0 {
		klog.Error(requestVar.VideoId)
		klog.Error(requestVar.Token)
		SendCommentListResponse(c, nil, errno.ParamErr)
		return
	}

	req := &useraction.CommentListRequest{VideoId: requestVar.VideoId, UserId: uId}
	cList, err := rpc.CommentList(context.Background(), req)
	if err != nil {
		SendCommentListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendCommentListResponse(c, cList, nil)
	return
}
