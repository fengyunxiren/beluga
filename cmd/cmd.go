package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "beluga",
	Short:        "beluga is a web server app",
	SilenceUsage: true,
	Long:         `beluga is a web server app`,
	Run: func(cmd *cobra.Command, args []string) {
		tips()
	},
}

func tips() {
	fmt.Println("Welcom to beluga")
	fmt.Println("Use beluga help to learn how to use")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
