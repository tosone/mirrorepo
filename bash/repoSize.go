package bash

import (
	"fmt"
	"os"
	"path/filepath"
)

func RepoSize(dir string) (size uint64, err error) {
	if isRepo := IsRepo(dir); !isRepo {
		err = fmt.Errorf("not a repo: %s", dir)
		return
	}
	err = filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += uint64(info.Size())
		}
		return err
	})
	return
}
