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

func (r *Router) SetRouter(url string, handler ...gin.HandlerFunc) {
	r.Engine.POST(url, handler...)
}
