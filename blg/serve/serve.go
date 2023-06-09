package serve

import (
	"blg/blg/ctrl"
	"blg/blg/mid"
	"blg/blg/router"
	"net/http"
	"sync"
	"sync/atomic"

	log "github.com/RRRRomeo/QLog/api"

	"github.com/gin-gonic/gin"
)

type BlgServe struct {
	Lock         sync.Mutex
	GlobalRouter *router.Router
	USERID       atomic.Int64
}

var Serve *BlgServe

func Init() bool {
	Serve = &BlgServe{
		GlobalRouter: router.NewRouter(gin.Default()),
	}
	return true
}

func GetBlgServe() *BlgServe {
	if Serve == nil {
		return nil
	}

	return Serve
}

func EventsHandler() bool {
	Serve.GlobalRouter.SetRouterPost("/api/register", mid.CheckUserExistMiddleware, ctrl.Register)
	Serve.GlobalRouter.SetRouterPost("/api/login", ctrl.Login)
	Serve.GlobalRouter.SetRouterPost("/api/create_post", mid.UserAuth(), ctrl.CreatePost)
	Serve.GlobalRouter.SetRouterPost("/api/update_post", mid.UserAuth(), ctrl.UpdatePost)
	return true
}

func Run(ipport string) bool {
	ser := &http.Server{
		Addr:    ipport,
		Handler: Serve.GlobalRouter.Engine,
	}

	err := ser.ListenAndServe()
	if err != nil {
		log.Debugf("Run serve fail:%s\n", err)
		return false
	}
	return true
}
