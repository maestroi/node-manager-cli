package setup

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func InstallDependencies() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Prefix = "Updating OS packages... "
	s.Start()
	runCommandSilently("apt-get", "update", "-y")
	s.Stop()
	color.Green("OS packages updated")

	s.Prefix = "Installing necessary dependencies... "
	s.Start()
	runCommandSilently("apt-get", "install", "-y", "software-properties-common")
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
