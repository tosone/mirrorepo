package bash

import (
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

func RepoAge(dir string) (age string, err error) {
	var isRepo bool
	if isRepo, err = IsRepo(dir); err != nil {
		return
	} else if !isRepo {
		return
	}
	// git log --reverse --pretty=oneline --format="%ar" | head -n 1
	cmd1 := exec.Command("git", "log", "--reverse", "--pretty=oneline", "--format='%ar'")
	cmd1.Dir = dir
	cmd2 := exec.Command("head", "-n", "1")
	cmd2.Stdin, err = cmd1.StdoutPipe()
	if err != nil {
		return
	}
	var stdoutPipe io.ReadCloser
	stdoutPipe, err = cmd2.StdoutPipe()
	if err != nil {
		return
	}
	err = cmd2.Start()
	if err != nil {
		return
	}
	err = cmd1.Start()
	if err != nil {
		return
	}
	stdoutBytes, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		return
	}
	err = cmd1.Wait()
	if err != nil {
		return
	}
	err = cmd2.Wait()
	if err != nil {
		return
	}
	age = strings.TrimSpace(string(stdoutBytes))
	return
}
