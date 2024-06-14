package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

func RunPlaybook(network, nodeType string) {
	fmt.Println("Running the Ansible playbook...")
	runCommand("ansible-playbook", "-i", "localhost,", "-c", "local", "/opt/nimiq-ansible/ansible/playbook.yml", "--extra-vars", fmt.Sprintf("network=%s node_type=%s", network, nodeType))
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

	// Copy the executable to the destination path
	input, err := os.ReadFile(exePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading executable: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(destPath, input, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing executable to /usr/local/bin: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Binary copied to /usr/local/bin")
}
