package core

import (
	"context"
)

type Interface interface {
	Set(ctx context.Context, path string, value string) error
	Get(ctx context.Context, path string) ([]byte, error)
	GetAll(ctx context.Context, path string) ([][]byte, error)
}
