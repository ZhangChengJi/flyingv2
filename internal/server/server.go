package server

import (
	"context"
	"flyingv2/conf"
	"flyingv2/logs"
	"flyingv2/pkg/etcd"
	"flyingv2/pkg/gin/plugin"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	systemAddress = 8080
)

type server struct {
	http *http.Server
}

type name interface {
}

func (s *server) init() error {
	logs.L.Info("Initialize etcd client...")
	err := s.setupEtcdClient()
	if err != nil {
		return err
	}
	logs.L.Info("Initialize server...")
	s.setupServer()
	return nil
}

func (s *server) Start(er chan error) {
	err := s.init()
	if err != nil {
		er <- err
		return
	}
	//start gin server
	logs.L.Info("start server Listening port: %s", zap.String("", s.http.Addr))
	go func() {
		err := s.http.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logs.L.Error("listen and serv fail: %s", zap.Error(err))
			er <- err
		}
	}()
}
func NewServer() *server {
	return &server{}
}

func (s *server) setupEtcdClient() error {
	if err := etcd.NewEtcdClient(conf.Endpoints); err != nil {
		logs.L.Error("init etcd client failed: %w", zap.Error(err))
		return err
	}
	return nil
}
func (s *server) setupServer() {
	address := fmt.Sprintf(":%d", systemAddress)
	r := plugin.Routers()
	s.http = &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    time.Duration(1000) * time.Millisecond,
		WriteTimeout:   time.Duration(5000) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *server) Stop() {
	s.shutdownServer(s.http)

}

func (s *server) shutdownServer(server *http.Server) {
	if server != nil {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logs.L.Error("Shutting down server error: %s", zap.Error(err))
		}
	}
}
