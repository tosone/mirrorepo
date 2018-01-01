package bash

import (
	"strconv"
	"strings"
)

func FileCount(dir string) (num uint64, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git ls-files | wc -l | tr -d ' '")
	num, err = strconv.ParseUint(strings.TrimSpace(string(stdout)), 10, 64)
	if err != nil {
		return
	}
	return
}
