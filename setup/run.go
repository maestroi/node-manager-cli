package setup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func RunPlaybook(network, nodeType, protocol string) {
	playbookPath := filepath.Join("/opt", fmt.Sprintf("%s-ansible", protocol), "ansible", "playbook.yml")
	color.Blue("Running the Ansible playbook...")
	runCommand("ansible-playbook", "-i", "localhost,", "-c", "local", playbookPath, "--extra-vars", fmt.Sprintf("network=%s node_type=%s", network, nodeType))
	color.Green("Ansible playbook run completed")
}

func RunPlaybookWithTags(network, nodeType, protocol, tags string) {
	playbookPath := filepath.Join("/opt", fmt.Sprintf("%s-ansible", protocol), "ansible", "playbook.yml")
	color.Blue("Running the Ansible playbook with tags: %s", tags)
	runCommand("ansible-playbook", "-i", "localhost,", "-c", "local", playbookPath, "--tags", tags, "--extra-vars", fmt.Sprintf("network=%s node_type=%s", network, nodeType))
	color.Green("Ansible playbook with tags %s run completed", tags)
}

func RunResetPlaybook(network, nodeType, protocol string) {
	RunPlaybookWithTags(network, nodeType, protocol, "reset")
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
