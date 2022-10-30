package cmd

import (
	"beluga/server"
	"beluga/server/common/database"

	"github.com/spf13/cobra"
)

var (
	migrateConfigPath string
	migrateCmd        = &cobra.Command{
		Use:          "migrate",
		Short:        "Migrate database",
		Example:      "beluga migrate -c etc/config.toml",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return server.SetUp(migrateConfigPath)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return database.MigrateDatabase(server.DBGenerator)
		},
	}
)

func init() {
	migrateCmd.PersistentFlags().StringVarP(&migrateConfigPath, "config", "c", "etc/config.toml", "Start server with provided configuration file")
	rootCmd.AddCommand(migrateCmd)
}
