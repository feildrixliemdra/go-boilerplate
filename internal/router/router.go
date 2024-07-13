package router

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	handler *gin.Engine
}

func NewRouter(handler *gin.Engine) Router {
	return &router{
		handler: handler,
	}
}

type Router interface {
	Init()
}

func (r *router) Init() {
	r.handler.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.handler.GET("", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
