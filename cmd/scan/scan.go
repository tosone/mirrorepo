package scan

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/bash"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskMgr"
	"github.com/tosone/mirrorepo/models"
	"github.com/tosone/mirrorepo/services/clone"
)

func Initialize(scanDir ...string) {
	var err error
	if err = models.Connect(); err != nil {
		logging.Fatal(err)
	}
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

			if strings.Index(p, ".git") != -1 {
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
				RealPlace: path.Join(viper.GetString("Setting.Repo"), base),
				Travel:    viper.GetInt("Setting.Travel"),
			}
			if err = repo.Create(); err != nil {
				logging.WithFields(logging.Fields{"repo": repo}).Error(err.Error())
				return err
			}
			err = taskMgr.Transport(taskMgr.ServiceCommand{
				Task:        "clone",
				Cmd:         "start",
				TaskContent: taskMgr.TaskContentClone{Repo: repo},
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
