package handlers

import (
	"context"
	"dy/cmd/api/kitex_gen/video"
	"dy/cmd/api/middleware"
	"dy/cmd/api/rpc"
	"dy/pkg/errno"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func PublishAction(c *gin.Context) {

	var requestVar PublishActionRequest
	mFile, err := c.FormFile("data")
	if err != nil {
		SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}
	if err := c.ShouldBind(&requestVar); err != nil {
		SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}

	file, err := mFile.Open()
	if err != nil {
		SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}
	data, err := ioutil.ReadAll(file)
	klog.Errorf("datalen", len(data))
	//先验证token
	uId, err := jwt.VerifyToken(requestVar.Token)
	klog.Errorf("uid", uId)
	if err != nil {
		SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}
	if uId <= 0 {
		klog.Error(uId)
		klog.Error(requestVar.Token)
		SendPublishActionResponse(c, errno.ParamErr)
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		klog.Error(err)
		SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}
	videoId := node.Generate().Int64()

	req := &video.PublishActionRequest{UserId: uId, Data: data, Title: requestVar.Title, VideoId: videoId}
	err = rpc.PublishAction(context.Background(), req)
	if err != nil {
		SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}
	SendPublishActionResponse(c, nil)
	return
}
