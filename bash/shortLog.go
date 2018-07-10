package bash

import (
	"regexp"
	"strings"
)

// ShortLog ..
func ShortLog(dir string) (authors map[string]string, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git log --all --format='%an <%ae>' --no-merges")
	if err != nil {
		return
	}
	var reg = regexp.MustCompile(`([\w\W]+)\s+<([\w\W]+)>`)

	authors = map[string]string{}
	for _, str := range strings.Split(strings.TrimSpace(string(stdout)), "\n") {
		matches := reg.FindStringSubmatch(str)
		if len(matches) == 3 {
			authors[matches[1]] = matches[2]
		}
	}
	return
}
