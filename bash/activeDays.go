package bash

import (
	"strconv"
	"strings"
)

func ActiveDays(dir string) (days uint64, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git ls-files | wc -l | tr -d ' '")
	days, err = strconv.ParseUint(strings.TrimSpace(string(stdout)), 10, 64)
	if err != nil {
		return
	}
	return
}
