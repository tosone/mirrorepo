package bash

import (
	"os/exec"

	"strings"

	"github.com/Unknwon/com"
)

// IsRepo ..
func IsRepo(dir string) (isRepo bool) {
	var stdout []byte
	var cmd *exec.Cmd

	if !com.IsDir(dir) {
		return
	}

	/* #nosec */
	cmd = exec.Command("/bin/sh", "-c", "git rev-parse --is-inside-git-dir")
	cmd.Dir = dir

	stdout, _ = cmd.Output()

	gitDir := strings.TrimSpace(string(stdout))

	/* #nosec */
	cmd = exec.Command("/bin/sh", "-c", "git rev-parse --is-inside-work-tree")
	cmd.Dir = dir

	stdout, _ = cmd.Output()

	gitWorkDir := strings.TrimSpace(string(stdout))

	return gitDir == "true" || gitWorkDir == "true"
}
