package ctrl

import (
	"blg/blg/model"
	"blg/blg/resp"
	"blg/types"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	usera, exists := c.Get("user")
	if !exists {
		resp.Fail(c, nil, "Failed to get user information from context")
		return
	}
	// 将用户信息转为 ReqUser 结构体
	user, ok := usera.(*model.User)
	if !ok {
		resp.Fail(c, nil, "Failed to convert user information")
		return
	}

	respCurrUser := types.RespGetCurrentUser{
		Account: user.Account,
		Nciname: user.Nickname,
		Avatar:  user.Avatar,
		Id:      user.Id,
	}

	resp.Success(c, gin.H{"currentUser": respCurrUser}, "get msg success!")

}
