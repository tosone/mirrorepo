package bash

import "strings"

func GetRemoteUrl(dir string) (url string, err error) {
	var stdout []byte
	stdout, err = Run(dir, "git remote get-url origin")
	if err != nil {
		return
	}
	url = strings.TrimSpace(string(stdout))
	return
}
