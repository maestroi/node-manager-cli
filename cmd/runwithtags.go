package cmd

import (
	"fmt"
	"node-manager-cli/setup"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var tags string

var runWithTagsCmd = &cobra.Command{
	Use:   "runwithtags",
	Short: "Run the Ansible playbook with custom tags",
	Run: func(cmd *cobra.Command, args []string) {
		runWithTags()
	},
}

func init() {
	rootCmd.AddCommand(runWithTagsCmd)
	runWithTagsCmd.Flags().StringVarP(&network, "network", "n", "testnet", "Network to run the playbook on")
	runWithTagsCmd.Flags().StringVarP(&nodeType, "node-type", "t", "validator", "Type of the node")
	runWithTagsCmd.Flags().StringVarP(&protocol, "protocol", "p", "nimiq", "Protocol to run (e.g., nimiq, another-protocol)")
	runWithTagsCmd.Flags().StringVarP(&path, "data-path", "d", "/opt", "location to install the datadir of node")
	runWithTagsCmd.Flags().StringVarP(&tags, "tags", "g", "", "Tags to use for the Ansible playbook")
}

func runWithTags() {
	if _, err := os.Stat(setup.ConfigFilePath); os.IsNotExist(err) {
		color.Red("Error: Configuration file does not exist. Please run setup first.")
		os.Exit(1)
	}

	config, err := setup.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config file: %v\n", err)
		os.Exit(1)
	}

	if tags == "" {
		color.Red("Error: You must specify tags to use for the Ansible playbook.")
		os.Exit(1)
	}

	color.Blue("Running Ansible playbook with tags %s for %s network with protocol %s", tags, config.Network, config.Protocol)
	setup.RunPlaybookWithTags(config.Network, config.NodeType, config.Protocol, config.DataPath, tags)
	color.Green("Ansible playbook with tags %s run completed", tags)
}
