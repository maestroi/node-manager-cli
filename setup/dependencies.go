package setup

import "fmt"

func InstallDependencies() {
	fmt.Println("Updating OS packages...")
	runCommandSilently("apt-get", "update", "-y")
	fmt.Println("Installing necessary dependencies...")
	runCommandSilently("apt-get", "install", "-y", "software-properties-common")
}

func InstallAnsible() {
	fmt.Println("Adding Ansible repository...")
	runCommandSilently("apt-add-repository", "--yes", "--update", "ppa:ansible/ansible")
	fmt.Println("Installing Ansible...")
	runCommandSilently("apt-get", "install", "-y", "ansible")
}

func InstallAnsibleGalaxyCollection() {
	fmt.Println("Installing Ansible Galaxy collection for Docker...")
	runCommand("ansible-galaxy", "collection", "install", "community.docker")
}
