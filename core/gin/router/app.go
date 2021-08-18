package router

import (
	"flyingv2/api"
	app2 "flyingv2/core/app"
	"flyingv2/core/constant"
	"flyingv2/core/factory"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
}

func (AppR *AppRouter) InitRouter(Router *gin.RouterGroup) {
	app := &api.AppApi{app2.App{factory.Create(constant.AppPrefix)}}
	{
		Router.GET("set", app.Set)
		Router.GET("list", app.List)
	}

}
