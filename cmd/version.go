package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Bind Cli Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bind Cli Version v1.0")
	},
}
