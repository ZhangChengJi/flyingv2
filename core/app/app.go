package app

import (
	"context"
	"flyingv2/core"
)

type App struct {
	I     core.Interface
	Name  string
	Group string
}

func (a *App) GetList() string {

	a.I.Set(context.Background(), "aa/1", "0000")
	return ""
}
