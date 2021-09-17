package plugin

import (
	_ "flyingv2/docs"
	"flyingv2/logs"
	"flyingv2/pkg/gin/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	logs.L.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logs.L.Info("register swagger handler")
	var d = router.RouterGroupApp.AppRouter
	public := Router.Group("")
	{
		d.InitRouter(public)
	}

	logs.L.Info("router register success")
	return Router

}
