package setup

import (
	"fmt"
	"node-manager-cli/config"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func UpdateRepository(protocol, branch string) {
	repoPath := AnsibleRepoPath(protocol)
	if localPath := os.Getenv("NIMIQ_ANSIBLE_PATH"); localPath != "" {
		installLocalAnsibleRepo(localPath, repoPath)
		return
	}

	repoURL, ok := config.ProtocolRepoMap[protocol]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Protocol '%s' is not supported\n", protocol)
		os.Exit(1)
	}
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

func installLocalAnsibleRepo(localPath, destPath string) {
	absLocalPath, err := filepath.Abs(localPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving local ansible path: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(absLocalPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: NIMIQ_ANSIBLE_PATH does not exist: %s\n", absLocalPath)
		os.Exit(1)
	}

	color.Blue("Installing Ansible repository from local path: %s", absLocalPath)
	if err := os.RemoveAll(destPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error removing existing ansible directory: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("cp", "-a", absLocalPath, destPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying local ansible repository: %v\n", err)
		os.Exit(1)
	}

	color.Green("Nimiq Ansible repository installed from local path")
}
