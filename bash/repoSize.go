package bash

import (
	"regexp"
	"strconv"
	"strings"
)

func RepoSize(dir string) (size uint64, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git count-objects -v")
	if err != nil {
		return
	}
	var reg = regexp.MustCompile(`size-pack:\s(\d+)`)
	for _, str := range strings.Split(string(stdout), "\n") {
		matches := reg.FindStringSubmatch(str)
		if len(matches) == 2 {
			size, err = strconv.ParseUint(matches[1], 10, 64)
			if err != nil {
				return
			}
		}
	}
	return
}
