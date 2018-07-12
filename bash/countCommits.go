package bash

import (
	"strconv"
	"strings"
)

// CountCommits ..
func CountCommits(dir string) (count uint64, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git rev-list --all --count")
	if err != nil {
		return
	}

	count, err = strconv.ParseUint(strings.TrimSpace(string(stdout)), 10, 64)
	if err != nil {
		return
	}
	return
}
