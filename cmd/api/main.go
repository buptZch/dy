package main

import (
	"context"
	"dy/cmd/api/kitex_gen/userbase"
	"github.com/gin-gonic/gin/binding"
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
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
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
			if err := c.ShouldBindBodyWith(&loginVar, binding.JSON); err != nil {
				klog.Errorf("====%+v\n", loginVar)
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
	v1 := r.Group("/douyin")
	user1 := v1.Group("/user")
	user1.POST("/login", authMiddleware.LoginHandler)
	user1.POST("/register", authMiddleware.RegisterHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "mee")
	})
	note1 := v1.Group("/note")
	note1.Use(authMiddleware.MiddlewareFunc())
	//note1.GET("/query", handlers.QueryNote)
	//note1.POST("", handlers.CreateNote)
	//note1.PUT("/:note_id", handlers.UpdateNote)
	//note1.DELETE("/:note_id", handlers.DeleteNote)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
