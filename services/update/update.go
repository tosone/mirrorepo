package update

import (
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/tosone/mirror-repo/bash"
	"github.com/tosone/mirror-repo/common/taskMgr"
	"github.com/tosone/mirror-repo/logging"
	"github.com/tosone/mirror-repo/models"
)

const serviceName = "update"

var updateLocker = new(sync.Mutex)

var updateList = map[uint]*models.Repo{}

func Initialize() {
	channel := make(chan taskMgr.ServiceCommand, 1)
	go func() {
		for {
			select {
			case control := <-channel:
				switch control.Cmd {
				case "start":
					for _, repo := range updateList {
						if control.TaskContent.(taskMgr.TaskContentClone).Repo.ID == repo.ID {
							return
						}
					}
					updateList[control.TaskContent.(taskMgr.TaskContentClone).Repo.ID] = control.TaskContent.(taskMgr.TaskContentClone).Repo
					updateLocker.Lock()
					update(control.TaskContent.(taskMgr.TaskContentUpdate))
					delete(updateList, control.TaskContent.(taskMgr.TaskContentClone).Repo.ID)
					updateLocker.Unlock()
				}
			}
		}
	}()
	taskMgr.Register(serviceName, channel)
}

// WaitAll ..
func WaitAll() {
	var done = make(chan bool)
	go func() {
		for {
			if len(updateList) == 0 {
				done <- true
				break
			}
			time.Sleep(time.Second)
		}
	}()
	<-done
}

func update(content taskMgr.TaskContentUpdate) {
	var err error
	var repo = content.Repo

	if !com.IsDir(repo.RealPlace) {
		logging.WithFields(logging.Fields{"repo": repo.RealPlace}).Error("Dir not found.")
	}

	defer func() {
		var status = "success"
		var msg = ""
		if err != nil {
			status = "error"
			msg = err.Error()
		}
		log := &models.Log{
			RepoID: repo.ID,
			Cmd:    serviceName,
			Status: status,
			Msg:    msg,
			Time:   time.Now(),
		}
		if err = log.Create(); err != nil {
			logging.Error(err.Error())
		}

		err = repo.UpdateByID()
		if err != nil {
			logging.Error(err.Error())
		}
	}()

	err = bash.Update(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}

	repo.LastTraveled = time.Now()
	detail(repo)
}

func detail(repo *models.Repo) {
	var err error

	repo.LastCommitCount = repo.CommitCount
	repo.CommitCount, err = bash.CountCommits(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}

	repo.LastSize = repo.Size
	repo.Size, err = bash.RepoSize(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}

	repo.LastCommitId = repo.CommitId
	repo.CommitId, err = bash.CommitId(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}
}
