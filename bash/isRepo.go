package bash

import (
	"strings"

	"errors"

	"github.com/Unknwon/com"
)

func IsRepo(dir string) (isRepo bool, err error) {
	if !com.IsDir(dir) {
		err = errors.New("dir is not exist")
		return
	}

	var stdout []byte
	stdout, err = Run(dir, "git rev-parse --is-inside-work-tree")
	if err != nil {
		return
	}

	return strings.TrimSpace(string(stdout)) == "true", err
}
