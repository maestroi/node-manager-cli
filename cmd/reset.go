package cmd

import (
	"fmt"
	"node-manager-cli/setup"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the Nimiq node ledger",
	Run: func(cmd *cobra.Command, args []string) {
		resetNode()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.Flags().StringVarP(&network, "network", "n", "testnet", "Network to reset the node on")
	resetCmd.Flags().StringVarP(&nodeType, "node-type", "t", "validator", "Type of the node")
	resetCmd.Flags().StringVarP(&protocol, "protocol", "p", "nimiq", "Protocol to reset (e.g., nimiq, another-protocol)")
}

func resetNode() {
	if _, err := os.Stat(setup.ConfigFilePath); os.IsNotExist(err) {
		color.Red("Error: Configuration file does not exist. Please run setup first.")
		os.Exit(1)
	}

	config, err := setup.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config file: %v\n", err)
		os.Exit(1)
	}

	color.Blue("Resetting Nimiq node ledger for %s network with protocol %s", config.Network, config.Protocol)
	setup.RunResetPlaybook(config.Network, config.NodeType, config.Protocol, config.DataPath)
	color.Green("Nimiq node ledger reset complete!")
}
