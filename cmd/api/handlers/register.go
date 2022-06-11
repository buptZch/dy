// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"context"
	"dy/cmd/api/kitex_gen/userbase"
	"dy/pkg/errno"
	"github.com/bwmarrin/snowflake"

	"dy/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// Register register user info
func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendUserBaseResponse(c, errno.ConvertErr(err), 0, "")
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendUserBaseResponse(c, errno.ParamErr, 0, "")
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		SendUserBaseResponse(c, errno.ConvertErr(err), 0, "")
		return
	}
	userid := node.Generate().Int64()

	err = rpc.CreateUser(context.Background(), &userbase.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
		UserId:   userid,
	})
	if err != nil {
		SendUserBaseResponse(c, errno.ConvertErr(err), 0, "")
		return
	}

	SendUserBaseResponse(c, errno.Success, userid, "")
}
