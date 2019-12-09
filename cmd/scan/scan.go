package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/bash"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
	"github.com/tosone/mirrorepo/services/clone"
)

// Initialize ..
func Initialize(scanDir ...string) {
	var err error

	for _, dir := range scanDir {
		var repoPreFix []string
		err = filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				logging.Panic(err.Error())
				return err
			}
			if !info.IsDir() {
				return nil
			}

			if base := filepath.Base(p); strings.HasPrefix(base, ".") {
				return nil
			}

			if strings.Contains(p, ".git") {
				return nil
			}

			for _, prefix := range repoPreFix {
				if strings.HasPrefix(p, prefix) {
					return nil
				}
			}

			var isRepo bool
			if isRepo = bash.IsRepo(p); !isRepo {
				return nil
			}
			repoPreFix = append(repoPreFix, p)
			var base = filepath.Base(p)

			var repo = &models.Repo{
				Address:   p,
				Status:    defination.Waiting,
				Name:      base,
				AliasName: uuid.NewV4().String(),
				Travel:    viper.GetInt("Setting.Travel"),
			}
			if err = repo.Create(); err != nil {
				logging.WithFields(logging.Fields{"repo": repo}).Error(err.Error())
				return err
			}
			err = taskmgr.Transport(taskmgr.ServiceCommand{
				Task:        "clone",
				Cmd:         "start",
				TaskContent: taskmgr.TaskContentClone{Repo: repo},
			})

			return err
		})
	}

	if err != nil {
		logging.Panic(err)
	}
	clone.WaitAll()

	fmt.Printf("\nScan ending.\n\n")
}
