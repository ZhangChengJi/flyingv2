package api

import (
	app2 "flyingv2/core/app"
	"flyingv2/core/etcd"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tsest struct {
}

func (t *Tsest) Test(c *gin.Context) {

	app := &app2.App{I: etcd.Storage}
	list := app.GetList()

	fmt.Println(list)
	c.String(http.StatusOK, list)
}
