package mid

import (
	"blg/blg/db"
	"blg/types"
	"net/http"

	qlog "github.com/RRRRomeo/QLog/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CheckUserExistMiddleware 是检查用户是否存在的中间件
func CheckUserExistMiddleware(c *gin.Context) {
	// return func(c *gin.Context) {
	var user types.ReqUser

	// 解析 JSON 请求体到 User 结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		// 解析错误，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	qlog.Debugf("loguser:%v\n", user)

	// 查找数据库
	var dbuser db.User
	if err := db.Global.Where(" email = ?", user.Email).First(&dbuser).Error; err == nil {
		qlog.Debugf("dbUser:%v\n", dbuser)
		// 用户已存在，返回冲突的错误响应
		c.JSON(http.StatusOK, gin.H{"error": "The user already exists"})
		c.Abort()
		return
	} else {
		if err == gorm.ErrRecordNotFound {
			// 用户不存在，继续处理下一个中间件或请求处理函数
			qlog.Debugf("user dont exist\n")
			c.Set("user", &user) // 将用户信息存储在上下文中
			c.Next()
			return
		}
		// 查询出错，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		c.Abort()
		return
	}

	// }
}
