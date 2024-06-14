package setup

import (
	"fmt"
	"node-manager-cli/config"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func UpdateRepository(protocol, branch string) {
	repoURL, ok := config.ProtocolRepoMap[protocol]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Protocol '%s' is not supported\n", protocol)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	if _, err := os.Stat("/opt/nimiq-ansible"); os.IsNotExist(err) {
		s.Prefix = "Cloning the Nimiq Ansible repository... "
		s.Start()
		if branch != "" {
			runCommandSilently("git", "clone", "-b", branch, repoURL, "/opt/nimiq-ansible")
		} else {
			runCommandSilently("git", "clone", repoURL, "/opt/nimiq-ansible")
		}
		s.Stop()
		color.Green("Nimiq Ansible repository cloned")
	} else {
		s.Prefix = "Updating the Nimiq Ansible repository... "
		s.Start()
		if branch != "" {
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "fetch")
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "checkout", branch)
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "pull")
		} else {
			runCommandSilently("git", "-C", "/opt/nimiq-ansible", "pull")
		}
		s.Stop()
		color.Green("Nimiq Ansible repository updated")
	}
}
