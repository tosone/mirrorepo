package bash

import "strings"

// GetRemoteURL ..
func GetRemoteURL(dir string) (url string, err error) {
  var stdout []byte
	stdout, err = Run(dir, "git remote get-url origin")
	if err != nil {
		return
	}
	url = strings.TrimSpace(string(stdout))
	return
}
