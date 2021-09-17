package core

import (
	"context"
	"flyingv2/internal/core/model"
)

type Interface interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (interface{}, error)
	List(ctx context.Context, ops *model.ListOptions) (*model.PageList, error)
}
