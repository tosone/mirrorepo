package bash

import "strings"

// CommitID ..
func CommitID(dir string) (commitID string, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git rev-parse HEAD")
	if err != nil {
		return
	}
	commitID = strings.TrimSpace(string(stdout))
	return
}
