package bash

import "strings"

// CommitId ..
func CommitId(dir string) (commitId string, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git rev-parse HEAD")
	if err != nil {
		return
	}
	commitId = strings.TrimSpace(string(stdout))
	return
}
