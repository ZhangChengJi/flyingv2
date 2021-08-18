package app

import (
	"context"
	"flyingv2/core"
)

type App struct {
	Api core.Interface
}

func (a *App) Set() string {

	a.Api.Set(context.Background(), "user", "1111")
	a.Api.Set(context.Background(), "dev", "1111")

	a.Api.Set(context.Background(), "ss", "1111")

	a.Api.Set(context.Background(), "ccc", "1111")

	a.Api.Set(context.Background(), "fffw", "1111")
	a.Api.Set(context.Background(), "12dsds", "1111")
	a.Api.Set(context.Background(), "2ffd", "1111")
	a.Api.Set(context.Background(), "efef", "1111")
	a.Api.Set(context.Background(), "fd", "1111")
	a.Api.Set(context.Background(), "ccsssc", "1111")
	a.Api.Set(context.Background(), "ddd", "1111")
	return ""

}
func (a *App) List() string {
	a.Api.List(context.Background(), "ddd	")
	//fmt.Println(string(b))
	return ""
}
