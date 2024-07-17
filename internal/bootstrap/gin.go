package bootstrap

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/handler"
	"go-boilerplate/internal/router"
	"reflect"

	"strings"
)

// use a single instance, it caches struct info
var (
	uni *ut.UniversalTranslator
)

func GetStructTag(f reflect.StructField, tag string) string {
	fieldTag := f.Tag.Get(tag)
	if tag == "" || fieldTag == "" {
		fieldTag = f.Tag.Get("json")
	}

	name := strings.SplitN(fieldTag, ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
func InitiateGinRouter(cfg *config.Config, handler *handler.Handler) *gin.Engine {
	r := gin.Default()

	//init router
	route := router.NewRouter(r, handler)
	route.Init()

	gin.SetMode(cfg.App.ReleaseMode)

	return r
}
