package router

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/handler"
)

type router struct {
	rtr     *gin.Engine
	handler *handler.Handler
}

func NewRouter(rtr *gin.Engine, handler *handler.Handler) Router {
	return &router{
		rtr,
		handler,
	}
}

type Router interface {
	Init()
}

func (r *router) Init() {

	r.rtr.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.rtr.GET("/example", r.handler.ExampleHandler.GetAll)
}
