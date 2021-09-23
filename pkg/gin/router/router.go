package router

import (
	"flyingv2/api/app"
	_ "flyingv2/docs"
	"flyingv2/internal/core"
	"flyingv2/logs"
	"flyingv2/pkg/gin/plugin"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.New()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	logs.L.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logs.L.Info("register swagger handler")
	p := &PrivateGroup{Router.Group("")}
	p.Register()
	logs.L.Info("router register success")
	return Router

}

type PublicGroup struct {
	*gin.RouterGroup
}
type PrivateGroup struct {
	*gin.RouterGroup
}

func (r *PrivateGroup) Register() {
	r.Use(plugin.JwtAuth())
	list := []core.RouteRegister{
		app.NewHandler(),
	}
	for _, register := range list {
		register.Router(r.RouterGroup)
	}
}
