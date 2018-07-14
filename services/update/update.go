package update

import (
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/bash"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
)

const serviceName = "update"

var updateLocker = new(sync.Mutex)

var updateList = map[uint]*models.Repo{}

// Initialize ..
func Initialize() {
	channel := make(chan taskmgr.ServiceCommand, 1)
	go func() {
		for control := range channel {
			switch control.Cmd {
			case "start":
				for _, repo := range updateList {
					if control.TaskContent.(taskmgr.TaskContentClone).Repo.ID == repo.ID {
						return
					}
				}
				updateList[control.TaskContent.(taskmgr.TaskContentClone).Repo.ID] = control.TaskContent.(taskmgr.TaskContentClone).Repo
				updateLocker.Lock()
				update(control.TaskContent.(taskmgr.TaskContentUpdate))
				delete(updateList, control.TaskContent.(taskmgr.TaskContentClone).Repo.ID)
				updateLocker.Unlock()
			}
		}
	}()
	taskmgr.Register(serviceName, channel)
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

func update(content taskmgr.TaskContentUpdate) {
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

	repo.LastCommitID = repo.CommitID
	repo.CommitID, err = bash.CommitID(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}
}
