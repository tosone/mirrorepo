package bash

import (
	"os/exec"
)

func Run(dir, script string) (stdout []byte, err error) {
	var isRepo bool
	if isRepo, err = IsRepo(dir); err != nil {
		return
	} else if !isRepo {
		return
	}

	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", script)
	cmd.Dir = dir

	stdout, err = cmd.Output()
	if err != nil {
		return
	}
	return
}
