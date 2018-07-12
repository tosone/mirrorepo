package bash

import (
	"fmt"
	"os/exec"
)

// Run ..
func Run(dir, script string) (stdout []byte, err error) {
	var isRepo bool
	if isRepo = IsRepo(dir); !isRepo {
		err = fmt.Errorf("not a repo: %s", dir)
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
