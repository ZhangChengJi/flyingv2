package rootcmd

import (
	"flyingv2/conf"
	"fmt"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootcmd.PersistentFlags().StringVarP(&conf.Endpoints, "etcd-endpoint", "c", "", "etcd-endpoint address")

}

var rootcmd = &cobra.Command{
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
		err := run()
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
