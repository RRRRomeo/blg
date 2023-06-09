package mid

import (
	"blg/blg/db"
	"blg/blg/model"
	"blg/tools/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		// 验证通过后获取claims中的userId
		userId := claims.UserId
		db := db.GetDB()
		user := &model.User{}
		db.First(user, userId)

		// 用户不存在
		if user.Id == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
