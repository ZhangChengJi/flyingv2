package router

import (
	"flyingv2/api/app"
	"flyingv2/api/system"
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
	b := &PublicGroup{Router.Group("")}
	b.Register()
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
	r.Use(plugin.JwtAuth()) //认证
	list := []core.RouteRegister{
		app.NewHandler(),
	}
	for _, register := range list {
		register.Router(r.RouterGroup)
	}
}
func (r *PublicGroup) Register() {
	list := []core.RouteRegister{
		system.NewHandler(),
	}
	for _, register := range list {
		register.Router(r.RouterGroup)
	}
}
