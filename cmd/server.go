package cmd

import (
	"beluga/server"

	"github.com/spf13/cobra"
)

var (
	configPath string
	serverCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start server",
		Example:      "beluga server -c etc/config.toml",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return server.SetUp(configPath)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.RunServer()
		},
	}
)

func init() {
	serverCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "etc/config.toml", "Start server with provided configuration file")
	rootCmd.AddCommand(serverCmd)
}
