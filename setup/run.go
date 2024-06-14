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
	fmt.Println("Copying binary to /usr/local/bin...")
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding executable path: %v\n", err)
		os.Exit(1)
	}
	destPath := filepath.Join("/usr/local/bin", "node-manager-cli")
	input, err := os.ReadFile(exePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading executable: %v\n", err)
		os.Exit(1)
	}
	if err = os.WriteFile(destPath, input, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing executable to /usr/local/bin: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Binary copied to /usr/local/bin")
}
