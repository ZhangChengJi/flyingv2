package app

import (
	"context"
	"flyingv2/core"
)

type App struct {
	Api core.Interface
	core.PageInfo
	*core.PageList
}

func (a *App) Set() string {

	a.Api.Set(context.Background(), "cc", "1111")
	a.Api.Set(context.Background(), "ces ", "1111")

	a.Api.Set(context.Background(), "cccc", "1111")

	a.Api.Set(context.Background(), "cccca", "1111")

	a.Api.Set(context.Background(), "cwcccc", "1111")
	a.Api.Set(context.Background(), "cxczz", "1111")
	a.Api.Set(context.Background(), "ccccccccc", "1111")
	a.Api.Set(context.Background(), "ccxcccc", "1111")
	a.Api.Set(context.Background(), "cqqqq", "1111")
	a.Api.Set(context.Background(), "csssss", "1111")
	a.Api.Set(context.Background(), "cuuu", "1111")
	return ""

}
func (a *App) List() error {
	opts := &core.ListOptions{
		PageInfo: a.PageInfo,
	}
	a.PageList = &core.PageList{}
	if err := a.Api.List(context.Background(), a.PageInfo.Key, opts, a.PageList); err != nil {
		return err
	}
	return nil
}
