package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tsest struct {
}

func (t *Tsest) Test(c *gin.Context) {
	c.String(http.StatusOK, "测试")
}
