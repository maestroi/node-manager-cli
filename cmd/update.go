package cmd

import (
	"bufio"
	"fmt"
	"node-manager-cli/setup"
	"node-manager-cli/utils"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var force bool

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Nimiq node to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		updateNode()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&force, "force", "f", false, "Force update even if the latest version is already installed")
	updateCmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch to use for the protocol repository (e.g., master, main)")
}

func updateNode() {
	config, err := setup.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config file: %v\n", err)
		os.Exit(1)
	}

	if branch == "" {
		branch = config.Branch
	}

	version, err := setup.GetVersion(config.Protocol, branch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching version: %v\n", err)
		os.Exit(1)
	}

	if config.Version != version || force {
		if config.Version != version {
			fmt.Printf("A new version (%s) is available. Current version is %s. Do you want to update? (y/n): ", version, config.Version)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			if strings.TrimSpace(strings.ToLower(input)) != "y" {
				fmt.Println("Update aborted.")
				return
			}
		}
		fmt.Println("Updating Nimiq node for", config.Network, "network with protocol", config.Protocol)
		setup.InstallDependencies()
		if !setup.IsCommandAvailable("ansible") {
			setup.InstallAnsible()
		}
		setup.InstallAnsibleGalaxyCollection()
		setup.UpdateRepository(config.Protocol, branch)
		config.Version = version
		config.Branch = branch
		setup.SaveConfig(config)
		setup.RunPlaybook(config.Network, config.NodeType, config.Protocol)
		fmt.Println("Nimiq node update complete!")

		ipAddress, err := utils.GetPublicIPAddress()
		if err != nil {
			color.Red("Error getting public IP address: %v", err)
		} else {
			color.Green("Grafana is available at: http://%s/grafana", ipAddress)
			color.Yellow("Default Grafana username: admin")
			color.Yellow("Default Grafana password: nimiq")
			color.Red("It is strongly recommended to change the default Grafana password.")
		}
	} else {
		fmt.Println("You already have the latest version.")
	}
}
