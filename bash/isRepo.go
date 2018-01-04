package bash

import (
	"os/exec"

	"strings"

	"github.com/Unknwon/com"
)

func IsRepo(dir string) (isRepo bool) {
	var stdout []byte
	var cmd *exec.Cmd

	if !com.IsDir(dir) {
		return
	}

	cmd = exec.Command("sh", "-c", "git rev-parse --is-inside-git-dir")
	cmd.Dir = dir

	stdout, _ = cmd.Output()

	gitDir := strings.TrimSpace(string(stdout))

	cmd = exec.Command("sh", "-c", "git rev-parse --is-inside-work-tree")
	cmd.Dir = dir

	stdout, _ = cmd.Output()

	gitWorkDir := strings.TrimSpace(string(stdout))

	return gitDir == "true" || gitWorkDir == "true"
}
