package cmd

import (
	"node-manager-cli/utils"

	"github.com/fatih/color"
)

func printGrafanaInfoIfEnabled(monitor bool) {
	if !monitor {
		return
	}

	ipAddress, err := utils.GetPublicIPAddress()
	if err != nil {
		color.Red("Error getting public IP address: %v", err)
		return
	}

	color.Green("Grafana is available at: http://%s/grafana", ipAddress)
	color.Yellow("Default Grafana username: admin")
	color.Yellow("Default Grafana password: nimiq")
	color.Red("It is strongly recommended to change the default Grafana password.")
}
