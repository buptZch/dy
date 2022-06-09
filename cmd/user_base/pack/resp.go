package pack

import (
	"errors"
	"time"

	"dy/pkg/errno"

	"dy/cmd/user_base/kitex_gen/userbase"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userbase.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *userbase.BaseResp {
	return &userbase.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
