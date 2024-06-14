package cmd

import (
	"fmt"
	"node-manager-cli/setup"
	"os"

	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup and remove all configurations and data files (use with CAUTION!!!)",
	Run: func(cmd *cobra.Command, args []string) {
		cleanupNode()
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

func cleanupNode() {
	fmt.Println("Cleaning up configuration and files...")
	os.RemoveAll("/opt/nimiq-ansible")
	os.Remove(setup.ConfigFilePath)
	os.Remove("/usr/local/bin/node-manager-cli")
	fmt.Println("Cleanup complete.")
}
