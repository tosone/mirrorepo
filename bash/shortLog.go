package bash

import (
	"fmt"
	"strings"
)

func ShortLog(dir string) (err error) {
	var stdout []byte
	stdout, err = Run(dir, "git log --all --format='%an <%ae>' --no-merges")
	if err != nil {
		return
	}

	fmt.Println(strings.Join(strings.Split(strings.TrimSpace(string(stdout)), "\n"), "@@@"))
	return
}
