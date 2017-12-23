package bash

import (
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"fmt"

	"github.com/Unknwon/com"
)

//git rev-list HEAD --count
func CountCommits(dir string) (count uint64, err error) {
	if !com.IsDir(dir) {
		return
	}
	cmd := exec.Command("git", "rev-list", "HEAD", "--count")
	//git shortlog -n -s $commit | LC_ALL=C awk  | column -t -s♪
	//cmd := exec.Command("git", "shortlog", "-n", "-s", "|", "sort")
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
		fmt.Println(err)
		return
	}
	count, err = strconv.ParseUint(strings.TrimSpace(string(stdoutBytes)), 10, 64)
	if err != nil {
		return
	}
	return
}
