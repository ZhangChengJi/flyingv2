package cmd

import (
	"flyingv2/conf"
	"flyingv2/internal/server"
	"flyingv2/logs"
	"fmt"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var command *cobra.Command

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
func init() {
	cobra.OnInitialize(func() {})
	command = NewFlyingCommand()

	command.PersistentFlags().StringVarP(&conf.Endpoints, "etcd-endpoint", "c", "", "etcd-endpoint address")

}

func NewFlyingCommand() *cobra.Command {

	return &cobra.Command{
		Use:   "flying",
		Short: "flying: Multi-environment cloud native distributed configuration center",
		Long: dedent.Dedent(`

				┌──────────────────────────────────────────────────────────┐
			    │ FlYING                                                   │
			    │ Cloud Native Distributed Configuration Center             │
			    │                                                          │
			    │ Please give us feedback at:                              │
			    │ https://github.com/ZhangChengJi/flyingv2/issues           │
			    └──────────────────────────────────────────────────────────┘
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args)
			err := startServer()
			if err != nil {
				return err
			}
			return nil
		},
	}
}
func startServer() error {
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

type serve interface {
	ListenAndServe() error
}
