package router

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/handler"
	"go-boilerplate/internal/middleware"
	"net/http"
)

type router struct {
	rtr     *gin.Engine
	handler *handler.Handler
	cfg     *config.Config
}

func NewRouter(rtr *gin.Engine, cfg *config.Config, handler *handler.Handler) Router {
	return &router{
		rtr,
		handler,
		cfg,
	}
}

type Router interface {
	Init()
}

func (r *router) Init() {
	// simulate some private data for basic auth
	var secrets = gin.H{
		"foo":  gin.H{"email": "foo@bar.com", "phone": "123433"},
		"test": gin.H{"email": "test@example.com", "phone": "666"},
	}

	r.rtr.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	exampleRouter := r.rtr.Group("/example")
	exampleRouter.GET("", r.handler.ExampleHandler.GetAll)
	exampleRouter.POST("", r.handler.ExampleHandler.Create)

	//example of JWT middleware
	authenticateHandler := r.rtr.Group("/authenticated")
	authenticateHandler.Use(middleware.JWTAuth(r.cfg.JWT.SecretKey))
	exampleRouter.GET("", r.handler.ExampleHandler.GetAll)

	authorized := r.rtr.Group("/secured", middleware.BasicAuth())

	// /admin/secrets endpoint
	// hit "localhost:8080/secured/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

}
