package setup

import (
	"fmt"
	"path/filepath"
)

func AnsibleRepoPath(protocol string) string {
	return filepath.Join("/opt", fmt.Sprintf("%s-ansible", protocol))
}

func AnsibleDir(protocol string) string {
	return filepath.Join(AnsibleRepoPath(protocol), "ansible")
}

func AnsiblePlaybookPath(protocol string) string {
	return filepath.Join(AnsibleDir(protocol), "playbook.yml")
}
