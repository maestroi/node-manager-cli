package cmd

import (
	"bufio"
	"fmt"
	"node-manager-cli/setup"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup and remove all configurations and files",
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
	fmt.Println("Configuration and files cleaned up.")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to stop and remove all Docker containers? (y/n): ")
	input, _ := reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(input)) == "y" {
		stopAndRemoveDockerContainers()
	} else {
		fmt.Println("Skipping Docker containers cleanup.")
	}

	fmt.Println("Cleanup complete.")
}

func stopAndRemoveDockerContainers() {
	color.Blue("Stopping all Docker containers...")
	containerIDs, err := getDockerContainerIDs()
	if err != nil {
		color.Red("Error fetching Docker container IDs: %v", err)
		return
	}

	if len(containerIDs) > 0 {
		if err := runCommand("docker", append([]string{"stop"}, containerIDs...)...); err != nil {
			color.Red("Error stopping Docker containers: %v", err)
		} else {
			color.Green("All Docker containers stopped.")
		}

		color.Blue("Removing all Docker containers...")
		if err := runCommand("docker", append([]string{"rm"}, containerIDs...)...); err != nil {
			color.Red("Error removing Docker containers: %v", err)
		} else {
			color.Green("All Docker containers removed.")
		}
	} else {
		color.Green("No Docker containers to stop or remove.")
	}
}

func getDockerContainerIDs() ([]string, error) {
	cmd := exec.Command("docker", "ps", "-aq")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	containerIDs := strings.Fields(string(output))
	return containerIDs, nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
