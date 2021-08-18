package api

import (
	"flyingv2/core/app"
	"github.com/gin-gonic/gin"
)

type AppApi struct {
	app.App
}

/**
 /registry/app/user-server {appId:user-server,name: 用户服务,group: [dev,test,pro]}
/registry/app/user-server
/registry/app/user-server

/registry/group/dev  {name: dev,txt:"测试库"}

/registry/group/app/dev {[]}

*/
//var a=app.App{factory.Create(constant.AppPrefix)}

func (app *AppApi) Set(*gin.Context) {
	app.App.Set()
}

func (app *AppApi) List(*gin.Context) {
	app.App.List()
}
