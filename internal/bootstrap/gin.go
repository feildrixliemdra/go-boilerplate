package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/router"
)

func RegistryGinRouter(releaseMode string) *gin.Engine {
	r := gin.Default()
	gin.SetMode(releaseMode)

	//init router
	route := router.NewRouter(r)
	route.Init()

	return r
}
