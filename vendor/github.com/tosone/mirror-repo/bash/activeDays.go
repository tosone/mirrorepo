package bash

import (
	"strconv"
	"strings"
)

func ActiveDays(dir string) (days uint64, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git log --pretty='format: %ai' | cut -d ' ' -f 2 | sort -r | uniq | awk '{ sum += 1 } END { print sum }'")
	days, err = strconv.ParseUint(strings.TrimSpace(string(stdout)), 10, 64)
	if err != nil {
		return
	}
	return
}
