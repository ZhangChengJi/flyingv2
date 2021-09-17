package router

import (
	"flyingv2/api"
	"flyingv2/conf"
	"flyingv2/pkg/etcd"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
}

func (AppR *AppRouter) InitRouter(Router *gin.RouterGroup) {
	app := new(api.AppApi)
	app.Store = etcd.New(conf.AppPrefix)
	appRouter := Router.Group("app")
	{
		appRouter.POST("set", app.Set)
		appRouter.GET("get", app.Get)
		appRouter.GET("list", app.List)
	}

}
