package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		Engine: engine,
	}
}

func (r *Router) SetRouterPost(url string, handler ...gin.HandlerFunc) {
	r.Engine.POST(url, handler...)
}

func (r *Router) SetRouterGet(url string, handler ...gin.HandlerFunc) {
	r.Engine.GET(url, handler...)
}
