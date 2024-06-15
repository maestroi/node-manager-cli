package cmd

import (
	"node-manager-cli/config"
	"node-manager-cli/setup"
	"node-manager-cli/utils"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var network string
var nodeType string
var protocol string
var branch string

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup or update Nimiq node",
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
		color.Red("Error: Configuration file already exists. A node is already set up on this system.")
		os.Exit(1)
	}

	if !config.IsValidConfiguration(protocol, network, nodeType) {
		color.Red("Error: Unsupported configuration. Use the 'list' command to see all supported configurations.")
		os.Exit(1)
	}

	color.Blue("Setting up %s node for %s network with node type %s", protocol, network, nodeType)
	setup.InstallDependencies()
	if !setup.IsCommandAvailable("ansible") {
		setup.InstallAnsible()
	}
	setup.InstallAnsibleGalaxyCollection()
	setup.UpdateRepository(protocol, branch)

	// Copy binary to /usr/local/bin, but handle in-use binary situation
	setup.CopyBinaryToUsrLocalBin()

	version, err := setup.GetVersion(protocol, branch)
	if err != nil {
		color.Red("Error fetching version: %v", err)
		os.Exit(1)
	}

	setup.SaveConfig(setup.Config{
		Protocol:   protocol,
		Network:    network,
		NodeType:   nodeType,
		Version:    version,
		Branch:     branch,
		CLIVersion: "1.0.0", // Set this to the current version of your CLI
	})
	setup.RunPlaybook(network, nodeType, protocol)
	color.Green("Nimiq node setup/update complete!")

	ipAddress, err := utils.GetPublicIPAddress()
	if err != nil {
		color.Red("Error getting public IP address: %v", err)
	} else {
		color.Green("Grafana is available at: http://%s/grafana", ipAddress)
		color.Yellow("Default Grafana username: admin")
		color.Yellow("Default Grafana password: nimiq")
		color.Red("It is strongly recommended to change the default Grafana password.")
	}
}
