package api

import (
	"context"
	"flyingv2/core"
	"flyingv2/core/app"
	"flyingv2/core/resp"
	"flyingv2/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppApi struct {
	Store core.Interface
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
func (app *AppApi) Get(c *gin.Context) {
	key, _ := c.Get("key")
	s := key.(string)
	if val, err := app.Store.Get(context.Background(), s); err != nil {
		return
	}
}

func (app *AppApi) List(c *gin.Context) {
	var query core.PageInfo
	_ = c.ShouldBind(&query)
	if list, err := app.Store.List(context.Background(), &core.ListOptions{PageInfo: query}); err != nil {
		logs.L.Error("get list", zap.Error(err))
		resp.FailWithMessage("get list failed", c)
	} else {
		resp.OkWithData(list, c)
	}

}
