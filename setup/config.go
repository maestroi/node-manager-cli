package setup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"node-manager-cli/config"
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
		return branch, nil
	}
	return getLatestVersion(protocol)
}

func getLatestVersion(protocol string) (string, error) {
	repoURL, ok := config.ProtocolRepoMap[protocol]
	if !ok {
		return "", fmt.Errorf("protocol '%s' is not supported", protocol)
	}
	parts := strings.Split(repoURL, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid repository URL: %s", repoURL)
	}
	owner := parts[len(parts)-2]
	repo := strings.TrimSuffix(parts[len(parts)-1], ".git")

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", owner, repo)
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tags []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return "", err
	}
	if len(tags) == 0 {
		return "", fmt.Errorf("no tags found in repository %s/%s", owner, repo)
	}
	return tags[0].Name, nil
}

func GetLatestCLIVersion() (string, error) {
	apiURL := "https://api.github.com/repos/maestroi/node-manager-cli/releases/latest"
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}
