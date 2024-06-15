package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetPublicIPAddress() (string, error) {
	response, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", fmt.Errorf("error getting public IP address: %v", err)
	}
	defer response.Body.Close()

	ip, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return strings.TrimSpace(string(ip)), nil
}
