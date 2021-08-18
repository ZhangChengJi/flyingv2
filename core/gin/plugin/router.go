package plugin

import (
	"flyingv2/core/gin/router"
	"flyingv2/logs"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	logs.L.Info("use middleware cors")
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logs.L.Info("register swagger handler")
	var d = router.RouterGroupApp.AppRouter
	public := Router.Group("")
	{
		d.InitRouter(public)
	}

	logs.L.Info("router register success")
	return Router

}
