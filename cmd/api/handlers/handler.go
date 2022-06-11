package handlers

import (
	"net/http"

	"dy/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

type UserBaseResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
}

// SendResponse pack response for login/register
func SendUserBaseResponse(c *gin.Context, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserBaseResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserId:  userId,
		Token:   token,
	})
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserParam struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}