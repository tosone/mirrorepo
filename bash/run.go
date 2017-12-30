package bash

import (
	"errors"
	"fmt"
	"os/exec"
)

func Run(dir, script string) (stdout []byte, err error) {
	var isRepo bool
	if isRepo, err = IsRepo(dir); err != nil {
		return
	} else if !isRepo {
		err = errors.New(fmt.Sprintf("not a repo: %s", dir))
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
