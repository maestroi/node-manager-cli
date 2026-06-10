package cmd

import (
	"node-manager-cli/setup"
	"node-manager-cli/utils"

	"github.com/fatih/color"
)

func playbookOptionsFromConfig(config setup.Config) setup.PlaybookOptions {
	return setup.PlaybookOptions{
		Monitor: config.Monitor,
		Homelab: config.Homelab,
	}
}

func printHomelabRPCInfo(homelab bool) {
	if !homelab {
		return
	}

	color.Green("Bare homelab install complete: node container only")
	ipAddress, err := utils.GetPublicIPAddress()
	if err != nil {
		color.Yellow("RPC is exposed on 0.0.0.0:8648. Use your host IP, for example: http://192.168.1.100:8648")
		return
	}

	color.Green("RPC is exposed at: http://%s:8648", ipAddress)
	color.Yellow("P2P is exposed on 0.0.0.0:8443")
}
