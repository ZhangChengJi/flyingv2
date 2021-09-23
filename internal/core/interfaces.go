package core

import (
	"context"
	"flyingv2/internal/core/model"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	List(ctx context.Context, ops *model.ListOptions) (*model.PageList, error)
	Update(ctx context.Context, key string, value string) error
}
type RouteHandler func() RouteRegister

type RouteRegister interface {
	Router(r *gin.RouterGroup)
}
