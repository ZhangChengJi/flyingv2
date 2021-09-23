package app

import (
	"context"
	"flyingv2/conf"
	"flyingv2/internal/core"
	"flyingv2/internal/core/model"
	"flyingv2/internal/core/resp"
	"flyingv2/logs"
	"flyingv2/pkg/etcd"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	Store core.Interface
}

func NewHandler() core.RouteRegister {
	return &Handler{
		etcd.New(conf.AppPrefix),
	}
}
func (h *Handler) Router(r *gin.RouterGroup) {
	g := r.Group("app")
	g.POST("set", h.Set)
	g.GET("get", h.Get)
	g.GET("list", h.List)
	g.PUT("update", h.Update)
}

// @Summary setapp
// @Description 进行app存储
// @Tags app
// @Accept mpfd
// @Produce json
// @Param object body model.App false "App"
// @Success 200 {string} json resp.Resp
// @Router /app/set [post]
func (h *Handler) Set(c *gin.Context) {
	var app model.App
	err := c.ShouldBind(&app)
	if err != nil {
		logs.L.Error("Parameter binding error", zap.Error(err))
		return
	}

	sj, _ := model.MarshalJSON(app)
	if err := h.Store.Set(context.Background(), app.AppId, string(sj)); err != nil {
		logs.L.Error("set app ", zap.Error(err))
		resp.FailWithMessage(fmt.Sprintf(err.Error(), "AppId", app.AppId), c)
	} else {
		resp.Ok(c)
	}
}

// @Summary 更新app
// @Description 更新app
// @Tags app
// @Accept mpfd
// @Produce json
// @Param object body model.App false "App"
//@Success 200 {object} resp.Resp
// @Router /app/update [put]
func (h *Handler) Update(c *gin.Context) {
	var ap model.App
	err := c.ShouldBind(&ap)
	if err != nil {
		logs.L.Error("Parameter binding error", zap.Error(err))
		return
	}
	sj, _ := model.MarshalJSON(ap)
	if err := h.Store.Update(context.Background(), ap.AppId, string(sj)); err != nil {
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
func (h *Handler) Get(c *gin.Context) {
	key := c.Query("key")
	if val, err := h.Store.Get(context.Background(), key); err != nil {
		logs.L.Error("get app key(%s)", zap.String("", key), zap.Error(err))
		resp.FailWithMessage("get app failed", c)
	} else {
		resp.OkWithData(val, c)
	}

}

func (h *Handler) List(c *gin.Context) {
	var query model.PageInfo
	_ = c.ShouldBind(&query)
	if list, err := h.Store.List(context.Background(), &model.ListOptions{PageInfo: query}); err != nil {
		logs.L.Error("get app list", zap.Error(err))
		resp.FailWithMessage("get app list failed", c)
	} else {
		resp.OkWithData(list, c)
	}

}
