package cmd

import (
	"fmt"
	"node-manager-cli/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of node-manager-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node-manager-cli version", config.CLIVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
