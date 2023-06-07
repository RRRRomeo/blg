package serve

import (
	"blg/blg/db"
	"blg/blg/mid"
	"blg/blg/router"
	"blg/types"
	"net/http"
	"sync/atomic"
	"time"

	log "github.com/RRRRomeo/QLog/api"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var GlobalServeEngine *gin.Engine
var GlobalRouter *router.Router
var USERID atomic.Int64

func Init() bool {
	if GlobalServeEngine == nil {
		GlobalServeEngine = gin.Default()
	}

	if GlobalRouter == nil {
		GlobalRouter = router.NewRouter(GlobalServeEngine)
	}

	// GlobalRouter.SetRouter("/register", Register)
	// GlobalServeEngine.Use(mid.CheckUserExistMiddleware())
	return true
}

func EventsHandler() bool {
	GlobalRouter.SetRouter("/register", mid.CheckUserExistMiddleware, Register)
	GlobalRouter.SetRouter("./login", Login)
	return true
}

func Run(ipport string) bool {
	err := GlobalServeEngine.Run(ipport)
	if err != nil {
		log.Debugf("Run serve fail:%s\n", err)
		return false
	}
	return true
}

func Login(ctx *gin.Context) {
	loginUser := &types.ReqUser{}
	dbUser := &db.User{}

	if err := ctx.ShouldBindJSON(loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ret := db.Global.Where("email = ?", loginUser.Email).First(dbUser)
	if ret.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": "user dont exist!"})
		return
	}
	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置声明（Payload）
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = dbUser.Username

	// 生成令牌字符串
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 返回令牌字符串
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func Register(ctx *gin.Context) {
	// var user types.ReqUser
	var dbuser db.User

	// 解析 JSON 请求体到 User 结构体
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	// 解析错误，返回错误响应
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// 从上下文中获取用户信息
	usera, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information from context"})
		return
	}
	// 将用户信息转为 ReqUser 结构体
	user, ok := usera.(*types.ReqUser)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert user information"})
		return
	}

	// 打印解析后的 User 对象
	log.Debugf("Name:%s\n", user.Name)
	log.Debugf("Email:%s\n", user.Email)

	// // 查找数据库
	// db.Global.Raw("SELECT * FROM t_user WHERE username = ?", user.Name).Scan(&dbuser)
	// if dbuser.Id != 0 {
	// 	ctx.JSON(http.StatusConflict, gin.H{"err": "the user is already exist"})
	// 	return
	// }

	if !genDbUser(&dbuser, user) {
		log.Errf("gen Db user fail!\n")
		return
	}
	ret := db.Global.Create(&dbuser)
	if ret.Error != nil {
		log.Errf("create user fail:%s\n", ret.Error)
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func genDbUser(dbu *db.User, requ *types.ReqUser) bool {
	dbu.Id = USERID.Add(1)
	dbu.Avatar = requ.Name
	dbu.Create_time = time.Now()
	dbu.Email = requ.Email
	// 生成密码的哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requ.Password), bcrypt.DefaultCost)
	if err != nil {
		// c.JSON(500, gin.H{"error": "Failed to generate hashed password"})
		return false
	}
	dbu.Password = string(hashedPassword)
	dbu.Nickname = requ.Name
	dbu.Last_login_time = time.Now()
	dbu.Username = requ.Name
	dbu.Update_time = time.Now()
	return true
}
