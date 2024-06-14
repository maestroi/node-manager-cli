package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"node-manager-cli/setup"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var autoupdateCmd = &cobra.Command{
	Use:   "autoupdate",
	Short: "Auto-update the node-manager-cli to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		autoUpdateCLI()
	},
}

func init() {
	rootCmd.AddCommand(autoupdateCmd)
}

func autoUpdateCLI() {
	config, err := setup.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config file: %v\n", err)
		os.Exit(1)
	}

	latestVersion, err := setup.GetLatestCLIVersion()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching latest CLI version: %v\n", err)
		os.Exit(1)
	}

	if config.CLIVersion != latestVersion {
		fmt.Printf("A new CLI version (%s) is available. Current version is %s. Do you want to update? (y/n): ", latestVersion, config.CLIVersion)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(input)) != "y" {
			fmt.Println("Update aborted.")
			return
		}

		fmt.Println("Updating CLI to the latest version...")
		tempFilePath := "/tmp/node-manager-cli"
		downloadURL := fmt.Sprintf("https://github.com/maestroi/node-manager-cli/releases/download/%s/node-manager-cli", latestVersion)
		err = downloadFile(tempFilePath, downloadURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error downloading the latest CLI version: %v\n", err)
			os.Exit(1)
		}
		os.Chmod(tempFilePath, 0755)

		err = moveFile(tempFilePath, "/usr/local/bin/node-manager-cli")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error moving the latest CLI version to final location: %v\n", err)
			os.Exit(1)
		}

		config.CLIVersion = latestVersion
		setup.SaveConfig(config)
		fmt.Println("CLI update complete!")
	} else {
		fmt.Println("You already have the latest CLI version.")
	}
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func moveFile(src, dst string) error {
	if err := os.Rename(src, dst); err != nil {
		if err := os.Remove(dst); err != nil {
			return err
		}
		if err := os.Rename(src, dst); err != nil {
			return err
		}
	}
	return nil
}
