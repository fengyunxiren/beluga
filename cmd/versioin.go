package cmd

import (
	"fmt"

	"beluga/global"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "beluga version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println("beluga version is ", global.VERSION)
	return nil
}
