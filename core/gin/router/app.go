package router

import (
	"flyingv2/api"
	"flyingv2/core/constant"
	"flyingv2/core/factory"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
}

func (AppR *AppRouter) InitRouter(Router *gin.RouterGroup) {
	app := new(api.AppApi)
	app.Store = factory.Create(constant.AppPrefix)
	//app2.App{factory.Create(constant.AppPrefix)}}
	appRouter := Router.Group("app")
	{
		appRouter.POST("set", app.Set)
		appRouter.GET("get", app.Get)
		appRouter.GET("list", app.List)
	}

}
