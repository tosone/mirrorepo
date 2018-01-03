package bash

import (
	"fmt"
)

func RemoteReset(dir, url string) (err error) {
	_, err = Run(dir, fmt.Sprintf("git remote set-url origin %s", url))
	return
}
