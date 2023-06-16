package mid

import (
	"blg/blg/db"
	"blg/blg/model"
	"blg/blg/resp"
	"blg/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CheckUserExistMiddleware 是检查用户是否存在的中间件
func CheckUserExistMiddleware(c *gin.Context) {
	// 使用 gin-cors 中间件来处理跨域请求
	var user types.ReqUser

	// 解析 JSON 请求体到 User 结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		// 解析错误，返回错误响应
		resp.Response(c, http.StatusBadRequest, 400, nil, err.Error())
		c.Abort()
		return
	}

	// 查找数据库
	var dbuser model.User
	if err := db.Global.Where(" account = ?", user.Account).First(&dbuser).Error; err == nil {
		resp.Fail(c, nil, "user already exist!")
		c.Abort()
		return
	} else {
		if err == gorm.ErrRecordNotFound {
			c.Set("user", &user) // 将用户信息存储在上下文中
			c.Next()
			return
		}
		// 查询出错，返回错误响应
		resp.Response(c, http.StatusInternalServerError, 500, nil, "fail to query db!")
		c.Abort()
		return
	}

}
