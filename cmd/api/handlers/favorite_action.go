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

func FavoriteAction(c *gin.Context) {

	var requestVar FavoriteActionRequest

	if err := c.ShouldBind(&requestVar); err != nil {
		SendFavoriteActionResponse(c, errno.ConvertErr(err))
		return
	}
	//先验证token
	uId, err := jwt.VerifyToken(requestVar.Token)
	klog.Errorf("uid", uId)
	if err != nil {
		SendFavoriteActionResponse(c, errno.ConvertErr(err))
		return
	}
	if uId <= 0 {
		klog.Error(uId)
		klog.Error(requestVar.Token)
		SendFavoriteActionResponse(c, errno.ParamErr)
		return
	}

	req := &useraction.FavoriteActionRequest{UserId: uId, VideoId: requestVar.VideoId, ActionType: requestVar.ActionType}
	err = rpc.FavoriteAction(context.Background(), req)
	if err != nil {
		SendFavoriteActionResponse(c, errno.ConvertErr(err))
		return
	}
	SendFavoriteActionResponse(c, nil)
	return
}
