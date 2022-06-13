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

func GetFavoriteList(c *gin.Context) {

	var requestVar GetFavoriteListRequest

	if err := c.BindQuery(&requestVar); err != nil {
		SendFavoriteListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	//先验证token
	_, err := jwt.VerifyToken(requestVar.Token)
	if err != nil {
		SendFavoriteListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	if requestVar.UserId <= 0 {
		klog.Error(requestVar.UserId)
		klog.Error(requestVar.Token)
		SendFavoriteListResponse(c, nil, errno.ParamErr)
		return
	}

	req := &useraction.FavoriteListRequest{UserId: requestVar.UserId}
	vList, err := rpc.FavoriteList(context.Background(), req)
	if err != nil {
		SendFavoriteListResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendFavoriteListResponse(c, vList, nil)
	return
}
