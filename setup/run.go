package setup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type PlaybookOptions struct {
	Monitor bool
	Homelab bool
}

func playbookExtraVars(network, nodeType, dataPath, nodeVersion string, opts PlaybookOptions) string {
	if opts.Homelab {
		return fmt.Sprintf(
			"network=%s node_type=%s data_path=%s node_version=%s homelab=true",
			network, nodeType, dataPath, nodeVersion,
		)
	}

	return fmt.Sprintf(
		"network=%s node_type=%s data_path=%s node_version=%s homelab=false monitor=%s nginx=true",
		network, nodeType, dataPath, nodeVersion, boolString(opts.Monitor),
	)
}

func boolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func RunPlaybook(network, nodeType, protocol, dataPath, nodeVersion string, opts PlaybookOptions) {
	playbookPath := AnsiblePlaybookPath(protocol)
	color.Blue("Running the Ansible playbook with node version %s...", nodeVersion)
	runCommand("ansible-playbook", "-i", "localhost,", "-c", "local", playbookPath, "--extra-vars", playbookExtraVars(network, nodeType, dataPath, nodeVersion, opts))
	color.Green("Ansible playbook run completed")
}

func RunPlaybookWithTags(network, nodeType, protocol, dataPath, tags, nodeVersion string, opts PlaybookOptions) {
	playbookPath := AnsiblePlaybookPath(protocol)
	color.Blue("Running the Ansible playbook with tags: %s (node version %s)", tags, nodeVersion)
	runCommand("ansible-playbook", "-i", "localhost,", "-c", "local", playbookPath, "--tags", tags, "--extra-vars", playbookExtraVars(network, nodeType, dataPath, nodeVersion, opts))
	color.Green("Ansible playbook with tags %s run completed", tags)
}

func RunResetPlaybook(network, nodeType, protocol, dataPath, nodeVersion string, opts PlaybookOptions) {
	RunPlaybookWithTags(network, nodeType, protocol, dataPath, "reset", nodeVersion, opts)
}

func CopyBinaryToUsrLocalBin() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding executable path: %v\n", err)
		os.Exit(1)
	}

	destPath := filepath.Join("/usr/local/bin", "node-manager-cli")

	// If the destination file already exists, rename it to a backup file
	if _, err := os.Stat(destPath); err == nil {
		backupPath := destPath + ".old"
		if err := os.Rename(destPath, backupPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error backing up the existing binary: %v\n", err)
			os.Exit(1)
		}
	}

	input, err := os.ReadFile(exePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading executable: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(destPath, input, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing executable to /usr/local/bin: %v\n", err)
		os.Exit(1)
	}
	color.Green("Binary copied to /usr/local/bin")
}
