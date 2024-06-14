package cmd

import (
	"fmt"
	"node-manager-cli/config"
	"node-manager-cli/setup"
	"os"

	"github.com/spf13/cobra"
)

var network string
var nodeType string
var protocol string
var branch string

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a protocol node.",
	Run: func(cmd *cobra.Command, args []string) {
		setupNode()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringVarP(&network, "network", "n", "testnet", "Network to deploy the node on")
	setupCmd.Flags().StringVarP(&nodeType, "node-type", "t", "validator", "Type of the node")
	setupCmd.Flags().StringVarP(&protocol, "protocol", "p", "nimiq", "Protocol to deploy (e.g., nimiq, another-protocol)")
	setupCmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch to use for the protocol repository (e.g., master, main)")
}

func setupNode() {
	if _, err := os.Stat(setup.ConfigFilePath); err == nil {
		fmt.Fprintf(os.Stderr, "Error: Configuration file already exists. A node is already set up on this system.\n")
		os.Exit(1)
	}

	if !config.IsValidConfiguration(protocol, network, nodeType) {
		fmt.Fprintf(os.Stderr, "Error: Unsupported configuration. Use the 'list' command to see all supported configurations.\n")
		os.Exit(1)
	}

	fmt.Println("Setting up Nimiq node for", network, "network with protocol", protocol)
	setup.InstallDependencies()
	if !setup.IsCommandAvailable("ansible") {
		setup.InstallAnsible()
	}
	setup.InstallAnsibleGalaxyCollection()
	setup.UpdateRepository(protocol, branch)

	version, err := setup.GetVersion(protocol, branch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching version: %v\n", err)
		os.Exit(1)
	}

	setup.SaveConfig(setup.Config{
		Protocol:   protocol,
		Network:    network,
		NodeType:   nodeType,
		Version:    version,
		Branch:     branch,
		CLIVersion: "0.1.0", // Set this to the current version of your CLI
	})
	setup.RunPlaybook(network, nodeType)
	// Copy binary to /usr/local/bin, but handle in-use binary situation
	setup.CopyBinaryToUsrLocalBin()
	fmt.Println("Nimiq node setup/update complete!")
}
