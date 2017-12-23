package bash

import (
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

func FileCount(dir string) (num uint64, err error) {
	var isRepo bool
	if isRepo, err = IsRepo(dir); err != nil {
		return
	} else if !isRepo {
		return
	}
	// git ls-files | wc -l | tr -d ' '
	cmd1 := exec.Command("git", "ls-files")
	cmd1.Dir = dir
	cmd2 := exec.Command("wc", "-l")
	cmd3 := exec.Command("tr", "-d", "' '")
	cmd2.Stdin, err = cmd1.StdoutPipe()
	if err != nil {
		return
	}
	cmd3.Stdin, err = cmd2.StdoutPipe()
	if err != nil {
		return
	}
	var stdoutPipe io.ReadCloser
	stdoutPipe, err = cmd3.StdoutPipe()
	if err != nil {
		return
	}

	err = cmd3.Start()
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
	err = cmd3.Wait()
	if err != nil {
		return
	}
	num, err = strconv.ParseUint(strings.TrimSpace(string(stdoutBytes)), 10, 64)
	if err != nil {
		return
	}
	return
}
