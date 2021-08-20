package api

import (
	"flyingv2/core"
	"flyingv2/core/app"
	"flyingv2/core/resp"
	"flyingv2/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func (app *AppApi) Set(c *gin.Context) {
	//c.ShouldBindJSON()
	app.App.Set()
}

func (app *AppApi) List(c *gin.Context) {
	var query core.PageInfo
	_ = c.ShouldBind(&query)
	app.App.PageInfo = query
	if err := app.App.List(); err != nil {
		logs.L.Error("get list", zap.Error(err))
		resp.FailWithMessage("get list failed", c)
	} else {
		resp.OkWithData(app.App.PageList, c)
	}

}
