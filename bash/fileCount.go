package bash

import (
	"strconv"
	"strings"
)

// FileCount ..
func FileCount(dir string) (num uint64, err error) {
	var stdout []byte
	if stdout, err = Run(dir, "git ls-files | wc -l | tr -d ' '"); err != nil {
		return
	}
	num, err = strconv.ParseUint(strings.TrimSpace(string(stdout)), 10, 64)

	return
}
