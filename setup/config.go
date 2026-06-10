package setup

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"node-manager-cli/config"
	"node-manager-cli/utils"
)

const ConfigFilePath = "/etc/node-manager-cli/node-manager-cli-config.json"

type Config struct {
	Protocol   string `json:"protocol"`
	Network    string `json:"network"`
	NodeType   string `json:"node_type"`
	Version    string `json:"version"`
	Branch     string `json:"branch,omitempty"`
	DataPath   string `json:"data_path"`
	CLIVersion string `json:"cli_version"`
}

func SaveConfig(config Config) {
	os.MkdirAll(filepath.Dir(ConfigFilePath), 0755)
	file, err := os.Create(ConfigFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding config file: %v\n", err)
		os.Exit(1)
	}
}

func LoadConfig() (Config, error) {
	file, err := os.Open(ConfigFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func GetVersion(protocol, branch string) (string, error) {
	if branch != "" {
		return normalizeNodeVersion(branch), nil
	}
	return getLatestNodeVersion(protocol)
}

func normalizeNodeVersion(version string) string {
	return strings.TrimPrefix(strings.TrimSpace(version), "v")
}

func getLatestNodeVersion(protocol string) (string, error) {
	repo, ok := config.ProtocolNodeReleaseMap[protocol]
	if !ok {
		return "", fmt.Errorf("protocol '%s' is not supported", protocol)
	}

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := utils.DecodeJSONFromURL(apiURL, &release); err != nil {
		return "", err
	}
	if release.TagName == "" {
		return "", fmt.Errorf("no latest release found for %s", repo)
	}
	return normalizeNodeVersion(release.TagName), nil
}

func GetLatestCLIVersion() (string, error) {
	apiURL := "https://api.github.com/repos/maestroi/node-manager-cli/releases/latest"
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := utils.DecodeJSONFromURL(apiURL, &release); err != nil {
		return "", err
	}
	if release.TagName == "" {
		return "", fmt.Errorf("no latest release found for node-manager-cli")
	}
	return release.TagName, nil
}
