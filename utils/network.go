package utils

import (
	"fmt"
	"strings"
)

func GetPublicIPAddress() (string, error) {
	ip, err := HTTPGetOK("https://api.ipify.org?format=text")
	if err != nil {
		return "", fmt.Errorf("error getting public IP address: %v", err)
	}

	return strings.TrimSpace(string(ip)), nil
}
