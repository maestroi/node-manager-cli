package setup

import (
	"fmt"
	"node-manager-cli/config"
	"os"
)

func UpdateRepository(protocol, branch string) {
	repoURL, ok := config.ProtocolRepoMap[protocol]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Protocol '%s' is not supported\n", protocol)
		os.Exit(1)
	}

	if _, err := os.Stat("/opt/nimiq-ansible"); os.IsNotExist(err) {
		fmt.Println("Cloning the Nimiq Ansible repository...")
		if branch != "" {
			runCommandSilently("git", "clone", "-b", branch, repoURL, "/opt/nimiq-ansible")
		} else {
			runCommandSilently("git", "clone", repoURL, "/opt/nimiq-ansible")
		}
	} else {
		fmt.Println("Updating the Nimiq Ansible repository...")
		if branch != "" {
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "fetch")
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "checkout", branch)
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "pull")
		} else {
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "pull")
		}
	}
}
