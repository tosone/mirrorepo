package travel

import (
	"log"
	"path"
	"time"

	"github.com/spf13/viper"
	"github.com/tosone/logging"

	"github.com/tosone/mirrorepo/bash"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
)

// Initialize ..
func Initialize() {
	for {
		var err error
		var repos *[]models.Repo
		if repos, err = new(models.Repo).GetAll(); err != nil {
			logging.Error(err)
		}
		for _, repo := range *repos {
			log.Println(repo)
			if repo.Status == defination.Error {
				if err = taskmgr.Transport(taskmgr.ServiceCommand{
					Task:        "clone",
					Cmd:         "start",
					TaskContent: taskmgr.TaskContentClone{Repo: &repo},
				}); err != nil {
					logging.Error(err)
				}
			} else if time.Since(repo.LastTraveled).Hours() > float64(repo.Travel) && repo.Status != defination.Waiting {
				var realPlace = path.Join(viper.GetString("Setting.Repo"), repo.AliasName)
				if err = bash.Update(realPlace); err != nil {
					logging.Error(err)
				}
			} else {
				logging.WithFields(logging.Fields{"repo": repo}).Debug("it not turn to pull the latest code")
			}
		}
		time.Sleep(time.Minute * 10)
	}
}
