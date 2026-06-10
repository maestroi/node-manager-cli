package setup

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func InstallDependencies() {
	RepairConflictingDockerAptSources()

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Updating OS packages... "
	s.Start()
	if err := runCommandCapture("apt-get", "update", "-y"); err != nil {
		s.Stop()
		color.Red("apt-get update failed. This is often caused by conflicting Docker apt sources from a previous install.")
		color.Yellow("Try: sudo rm -f /etc/apt/sources.list.d/docker.list && sudo apt-get update")
		os.Exit(1)
	}
	s.Stop()
	color.Green("OS packages updated")

	s.Prefix = "Installing necessary dependencies... "
	s.Start()
	if err := runCommandCapture("apt-get", "install", "-y", "software-properties-common"); err != nil {
		s.Stop()
		os.Exit(1)
	}
	s.Stop()
	color.Green("Necessary dependencies installed")
}

func InstallAnsible() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Adding Ansible repository... "
	s.Start()
	runCommandSilently("apt-add-repository", "--yes", "--update", "ppa:ansible/ansible")
	s.Stop()
	color.Green("Ansible repository added")

	s.Prefix = "Installing Ansible... "
	s.Start()
	runCommandSilently("apt-get", "install", "-y", "ansible")
	s.Stop()
	color.Green("Ansible installed")
}

func InstallAnsibleGalaxyCollection() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Installing Ansible Galaxy collection for Docker... "
	s.Start()
	runCommand("ansible-galaxy", "collection", "install", "community.docker")
	s.Stop()
	color.Green("Ansible Galaxy collection for Docker installed")
}
