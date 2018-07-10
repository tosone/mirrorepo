package travel

import (
	"time"

	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskMgr"
	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/models"
)

func Initialize() {

}

func travel() {
	var err error
	var repos []*models.Repo
	repos, err = new(models.Repo).GetAll()
	if err != nil {
		logging.Error(err.Error())
	}
	for _, repo := range repos {
		if repo.Status == defination.Error {
			err = taskMgr.Transport(taskMgr.ServiceCommand{
				Task:        "clone",
				Cmd:         "start",
				TaskContent: taskMgr.TaskContentClone{Repo: repo},
			})
			if err != nil {
				logging.Error(err.Error())
			}
		} else if time.Since(repo.LastTraveled).Hours() > float64(repo.Travel) && repo.Status != defination.Waiting {
			err = taskMgr.Transport(taskMgr.ServiceCommand{
				Task:        "update",
				Cmd:         "start",
				TaskContent: taskMgr.TaskContentClone{Repo: repo},
			})
			if err != nil {
				logging.Error(err.Error())
			}
		} else {
			logging.WithFields(logging.Fields{"repo": repo}).Debug("It not turn to pull the latest code.")
		}
	}
	time.Sleep(time.Minute * 10)
}
