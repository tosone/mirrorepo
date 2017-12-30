package bash

import (
	"fmt"
	"os/exec"
	"strings"
)

func ShortLoooooooog(dir string) (err error) {
	var cmd *exec.Cmd
	var stdoutBytes []byte
	cmd = exec.Command("sh", "-c", "git log --all --format='%an <%ae>' --no-merges")
	cmd.Dir = dir
	stdoutBytes, err = cmd.Output()

	fmt.Println(strings.Join(strings.Split(strings.TrimSpace(string(stdoutBytes)), "\n"), "@@@"))
	return
}
