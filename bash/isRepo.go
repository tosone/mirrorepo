package bash

import (
	"io"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/Unknwon/com"
)

func IsRepo(dir string) (isRepo bool, err error) {
	if !com.IsDir(dir) {
		return
	}
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree") // git rev-parse --is-inside-work-tree
	cmd.Dir = dir
	var stdoutPipe io.ReadCloser
	var stdoutBytes []byte

	stdoutPipe, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	stdoutBytes, err = ioutil.ReadAll(stdoutPipe)
	if err != nil {
		return
	}
	err = cmd.Wait()
	if err != nil {
		return
	}

	return strings.TrimSpace(string(stdoutBytes)) == "true", err
}
