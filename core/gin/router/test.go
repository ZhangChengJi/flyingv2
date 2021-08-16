package router

import (
	"flyingv2/api"
	"github.com/gin-gonic/gin"
)

type DD struct {
}

func (s *DD) InitRouter(Router *gin.RouterGroup) {
	var s2 = new(api.Tsest)
	{
		Router.GET("test", s2.Test)
	}

}
