package main

import (
	"context"
	"dy/cmd/api/kitex_gen/userbase"
	"dy/pkg/errno"
	"net/http"
	"time"

	"dy/cmd/api/handlers"
	"dy/cmd/api/middleware"
	"dy/cmd/api/rpc"
	"dy/pkg/constants"
	"dy/pkg/tracer"
	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour * 10000,
		MaxRefresh: time.Hour * 10000,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			//klog.Errorf("xxx", c.Get("user_name"))
			//if err := c.ShouldBindBodyWith(&loginVar, binding.JSON); err != nil {
			//	klog.Errorf("====%+v\n", loginVar)
			//	return "", jwt.ErrMissingLoginValues
			//}
			loginVar.UserName = c.Query("username")
			loginVar.PassWord = c.Query("password")
			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				klog.Error(errno.ParamErr)
				return "", jwt.ErrMissingLoginValues
			}
			klog.Errorf("====%+v\n", loginVar)
			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {

				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &userbase.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	r.Use(gin.Logger())
	r.POST("/douyin/user/login/", authMiddleware.LoginHandler)
	r.POST("/douyin/user/register/", authMiddleware.RegisterHandler)
	r.GET("/douyin/user/", handlers.GetUser)
	v1 := r.Group("/douyin")
	publish1 := v1.Group("/publish")
	publish1.GET("/list/", handlers.GetPublishVideo)
	publish1.POST("/action/", handlers.PublishAction)

	favorite1 := v1.Group("/favorite")
	favorite1.POST("/action/", handlers.FavoriteAction)
	favorite1.GET("/list/", handlers.GetFavoriteList)

	feed1 := v1.Group("/feed")
	feed1.GET("/", handlers.GetFeed)

	comment1 := v1.Group("/comment")
	comment1.POST("/action/", handlers.CommentAction)
	comment1.GET("/list/", handlers.GetCommentList)

	relation1 := v1.Group("relation")
	relation1.POST("/action/", handlers.RelationAction)
	relation1.GET("/follow/list/", handlers.GetFollowList)
	relation1.GET("/follower/list/", handlers.GetFollowerList)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
