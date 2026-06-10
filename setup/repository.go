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

	repoPath := AnsibleRepoPath(protocol)
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		s.Prefix = "Cloning the Nimiq Ansible repository... "
		s.Start()
		if branch != "" {
			runCommandSilently("git", "clone", "-b", branch, repoURL, repoPath)
		} else {
			runCommandSilently("git", "clone", repoURL, repoPath)
		}
		s.Stop()
		color.Green("Nimiq Ansible repository cloned")
	} else {
		s.Prefix = "Updating the Nimiq Ansible repository... "
		s.Start()
		if branch != "" {
			runCommandSilently("git", "-C", repoPath, "fetch")
			runCommandSilently("git", "-C", repoPath, "checkout", branch)
			runCommandSilently("git", "-C", repoPath, "pull")
		} else {
			runCommandSilently("git", "-C", repoPath, "pull")
		}
		s.Stop()
		color.Green("Nimiq Ansible repository updated")
	}
}
