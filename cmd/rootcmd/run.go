package rootcmd

import (
	"flyingv2/internal/server"
	"flyingv2/logs"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func run() error {
	logs.NewLog()
	logs.L.Info("Initialize flying server...")
	serve := server.NewServer()
	er := make(chan error, 3)
	serve.Start(er)

	// Signal received to the process externally.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		logs.L.Info("flying  server receive %s and start shutting down", zap.String("", sig.String()))
		serve.Stop()
		logs.L.Info("flying server exited")
	case err := <-er:
		logs.L.Error("flying server start failed: %s", zap.String("", err.Error()))
		return err
	}
	return nil

}
