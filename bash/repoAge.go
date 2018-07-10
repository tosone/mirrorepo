package bash

import (
	"strings"
)

// RepoAge ..
func RepoAge(dir string) (age string, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git log --reverse --pretty=oneline --format=\"%ar\" | head -n 1")
	if err != nil {
		return
	}
	age = strings.TrimSpace(string(stdout))
	return
}
