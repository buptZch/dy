package handlers

import (
	"context"
	"dy/cmd/api/kitex_gen/useraction"
	"dy/cmd/api/middleware"
	"dy/cmd/api/rpc"
	"dy/pkg/errno"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {

	var requestVar CommentActionRequest

	if err := c.ShouldBind(&requestVar); err != nil {
		SendCommentActionResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	uId, err := jwt.VerifyToken(requestVar.Token)
	klog.Errorf("uid", uId)
	if err != nil {
		SendCommentActionResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if uId <= 0 {
		klog.Error(uId)
		klog.Error(requestVar.Token)
		SendCommentActionResponse(c, nil, errno.ParamErr)
		return
	}

	if requestVar.ActionType == 1 {
		node, err := snowflake.NewNode(1)
		if err != nil {
			klog.Error(err)
			SendUserBaseResponse(c, errno.ConvertErr(err), 0, "")
			return
		}
		requestVar.CommentId = node.Generate().Int64()
	}
	req := &useraction.CommentActionRequest{UserId: uId, VideoId: requestVar.VideoId,
		ActionType: requestVar.ActionType, CommentId: requestVar.CommentId, CommentText: requestVar.CommentText}
	comment, err := rpc.CommentAction(context.Background(), req)
	if err != nil {
		SendCommentActionResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendCommentActionResponse(c, comment, nil)
	return
}
