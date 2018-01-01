package travel

import (
	"time"

	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/taskMgr"
	"github.com/tosone/mirror-repo/logging"
	"github.com/tosone/mirror-repo/models"
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
		if time.Since(repo.LastTraveled).Hours() > float64(repo.Travel) && repo.Status != defination.Waiting {
			err = taskMgr.Transport(taskMgr.ServiceCommand{
				Task:        "clone",
				Cmd:         "start",
				TaskContent: taskMgr.TaskContentClone{Repo: repo},
			})

			if err != nil {
				logging.Error(err.Error())
			}
		}
	}
	time.Sleep(time.Minute * 10)
}
