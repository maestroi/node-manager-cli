package cmd

import (
	"fmt"
	"node-manager-cli/config"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all supported protocols, networks, and node types",
	Run: func(cmd *cobra.Command, args []string) {
		listSupportedConfigurations()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listSupportedConfigurations() {
	fmt.Println("Supported Protocols, Networks, and Node Types:")
	for protocol, networks := range config.SupportedConfigurations {
		fmt.Printf("Protocol: %s\n", protocol)
		for network, nodeTypes := range networks {
			fmt.Printf("  Network: %s\n", network)
			fmt.Printf("    Node Types: %s\n", nodeTypes)
		}
	}
}
