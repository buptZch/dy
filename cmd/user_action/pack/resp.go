package pack

import (
	"errors"
	"time"

	"dy/pkg/errno"

	"dy/cmd/user_action/kitex_gen/useraction"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *useraction.BaseResp {
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

func baseResp(err errno.ErrNo) *useraction.BaseResp {
	return &useraction.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
