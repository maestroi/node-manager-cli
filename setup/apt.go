package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func RepairConflictingDockerAptSources() {
	sourcesDir := "/etc/apt/sources.list.d"
	entries, err := os.ReadDir(sourcesDir)
	if err != nil {
		return
	}

	var unsignedDockerFiles []string
	hasSignedByDockerSource := false

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "docker") {
			continue
		}

		path := filepath.Join(sourcesDir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		if !strings.Contains(string(content), "download.docker.com/linux/ubuntu") {
			continue
		}

		if strings.Contains(string(content), "signed-by=") {
			hasSignedByDockerSource = true
			continue
		}

		unsignedDockerFiles = append(unsignedDockerFiles, path)
	}

	if !hasSignedByDockerSource || len(unsignedDockerFiles) == 0 {
		return
	}

	for _, path := range unsignedDockerFiles {
		if err := os.Remove(path); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to remove conflicting Docker apt source %s: %v\n", path, err)
			continue
		}
		color.Yellow("Removed conflicting Docker apt source: %s", path)
	}
}
