package api

import (
	"context"
	"encoding/json"
	"flyingv2/internal/core"
	"flyingv2/internal/core/model"
	"flyingv2/internal/core/resp"
	"flyingv2/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppApi struct {
	Store core.Interface
}

// @Summary setapp
// @Description 进行app存储
// @Tags app
// @Accept mpfd
// @Produce json
// @Param object query core.App false "App"
// @Router /app/set [post]
func (app *AppApi) Set(c *gin.Context) {
	var ap model.App
	err := c.ShouldBind(&ap)
	if err != nil {
		logs.L.Error("Parameter binding error", zap.Error(err))
		return
	}
	sj, _ := json.Marshal(ap)
	if err := app.Store.Set(context.Background(), ap.AppId, string(sj)); err != nil {
		logs.L.Error("set app ", zap.Error(err))
		resp.FailWithMessage("set app failed", c)
	} else {
		resp.Ok(c)
	}
}

// @Summary 获取app
// @Description 根据key进行获取app
// @Tags app
// @Accept mpfd
// @Produce json
// @Param key query string true "Key"
//@Success 200 {object} resp.Resp
// @Router /app/get [get]
func (app *AppApi) Get(c *gin.Context) {
	key := c.Query("key")
	if val, err := app.Store.Get(context.Background(), key); err != nil {
		logs.L.Error("get app key(%s)", zap.String("", key), zap.Error(err))
		resp.FailWithMessage("get app failed", c)
	} else {
		resp.OkWithData(val, c)
	}

}

func (app *AppApi) List(c *gin.Context) {
	var query model.PageInfo
	_ = c.ShouldBind(&query)
	if list, err := app.Store.List(context.Background(), &model.ListOptions{PageInfo: query}); err != nil {
		logs.L.Error("get app list", zap.Error(err))
		resp.FailWithMessage("get app list failed", c)
	} else {
		resp.OkWithData(list, c)
	}

}
