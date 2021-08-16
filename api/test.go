package api

import (
	"context"
	"flyingv2/core/etcd"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tsest struct {
}

func (t *Tsest) Test(c *gin.Context) {

	re, _ := aa.Get(context.Background(), "", false)
	fmt.Println(re)
	c.String(http.StatusOK, "")
}
