package coremain

import (
	"flyingv2/core/etcd"
	"flyingv2/core/factory"
	plugin "flyingv2/core/gin/plugin"
	"flyingv2/logs"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func init() {
	logs.NewLog()
	etcd.Storage, _ = factory.Create()

}

func Run() {
	// start endpoints server
	server()
}

const (
	systemAddress = 8080
)

func server() {
	address := fmt.Sprintf(":%d", systemAddress)
	router := plugin.Routers()
	s := initServer(address, router)
	time.Sleep(10 * time.Microsecond)
	logs.L.Info("server run success on ", zap.String("address", address))
	logs.L.Error(s.ListenAndServe().Error())

}

type serve interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) serve {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
