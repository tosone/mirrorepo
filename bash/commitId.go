package bash

import "strings"

func CommitId(dir string) (commitId string, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git rev-parse HEAD")
	if err != nil {
		return
	}
	commitId = strings.TrimSpace(string(stdout))
	return
}
