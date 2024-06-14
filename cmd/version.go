package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0.0" // Update this as needed

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of node-manager-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node-manager-cli version", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
