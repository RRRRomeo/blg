package ctrl

import (
	"blg/blg/db"
	"blg/blg/dto"
	"blg/blg/model"
	"blg/blg/resp"
	"blg/tools/common"
	"blg/types"
	"net/http"

	log "github.com/RRRRomeo/QLog/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	loginUser := &types.ReqUser{}
	dbUser := &model.User{}

	if err := ctx.ShouldBindJSON(loginUser); err != nil {
		resp.Fail(ctx, nil, err.Error())
		return
	}

	ret := db.Global.Where("account = ?", loginUser.Account).First(dbUser)
	if ret.Error != nil {
		resp.Fail(ctx, nil, "user dont exist:"+ret.Error.Error())
		return
	}
	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		resp.Response(ctx, http.StatusUnauthorized, 401, nil, "username or password fail:"+err.Error())
		return
	}

	tokenString, err := common.ReleaseToken(*dbUser)
	if err != nil {
		resp.Response(ctx, http.StatusInternalServerError, 500, nil, "release token fail:"+err.Error())
		return
	}

	// 返回令牌字符串
	resp.Success(ctx, gin.H{"token": tokenString}, "handle suecces!")

}

func Register(ctx *gin.Context) {
	// var user types.ReqUser
	var dbuser model.User

	usera, exists := ctx.Get("user")
	if !exists {
		resp.Fail(ctx, nil, "Failed to get user information from context")
		return
	}
	// 将用户信息转为 ReqUser 结构体
	user, ok := usera.(*types.ReqUser)
	if !ok {
		resp.Fail(ctx, nil, "Failed to convert user information")
		return
	}

	// 打印解析后的 User 对象
	// log.Debugf("Name:%s\n", user.Name)
	// log.Debugf("Email:%s\n", user.Email)

	if !dto.UserDto(&dbuser, user) {
		log.Errf("gen Db user fail!\n")
		resp.Response(ctx, http.StatusInternalServerError, 500, nil, "dto fail!")
		return
	}

	ret := db.Global.Create(&dbuser)
	if ret.Error != nil {
		log.Errf("create user fail:%s\n", ret.Error)
		resp.Response(ctx, http.StatusInternalServerError, 500, nil, "create data in db fail:"+ret.Error.Error())
		return
	}

	rspuser := types.RespGetCurrentUser{
		Id:       dbuser.Id,
		Account:  user.Account,
		Nickname: user.NickName,
		Avatar:   "https://s1.ax1x.com/2023/06/25/pCNLtdP.png",
	}
	// 返回成功响应
	resp.Success(ctx, gin.H{"data": rspuser}, "register success!")
	// ctx.JSON(http.StatusOK, gin.H{'Oauth-Token': "token",})
}
