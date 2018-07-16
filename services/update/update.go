package update

import (
	"path"
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/spf13/viper"
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
			var task = control.TaskContent.(taskmgr.TaskContentUpdate)
			switch control.Cmd {
			case "start":
				for _, repo := range updateList {
					if task.Repo.ID == repo.ID {
						return
					}
				}
				updateList[task.Repo.ID] = task.Repo
				updateLocker.Lock()
				if err := update(task); err != nil {
					logging.Error(err)
				}
				delete(updateList, task.Repo.ID)
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

func update(content taskmgr.TaskContentUpdate) (err error) {
	var repo = content.Repo
	var realPlace = path.Join(viper.GetString("Setting.Repo"), repo.AliasName)

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
		}
		if err = log.Create(); err != nil {
			logging.Error(err.Error())
		}

		err = repo.UpdateByID()
		if err != nil {
			logging.Error(err.Error())
		}
	}()

	if !com.IsDir(realPlace) {
		logging.WithFields(logging.Fields{"repo": realPlace}).Error("dir not found")
	}

	err = bash.Update(realPlace)
	if err != nil {
		logging.Error(err.Error())
	}

	repo.LastTraveled = time.Now()
	if repo.CommitCount, err = bash.CountCommits(realPlace); err != nil {
		return
	}

	{
		var commit string
		var size uint64
		if size, err = bash.RepoSize(realPlace); err != nil {
			return
		}
		if commit, err = bash.CommitID(realPlace); err != nil {
			return
		}
		var historyInfo = &models.HistoryInfo{RepoLastID: repo.HistoryInfoID}
		var oldHistoryInfo = new(models.HistoryInfo)
		if oldHistoryInfo, err = historyInfo.Find(); err != nil {
			return
		}
		if oldHistoryInfo.Commit == commit {
			return
		}
		historyInfo.Size = size
		historyInfo.Commit = commit
		historyInfo.RepoID = repo.ID
		if err = historyInfo.Create(); err != nil {
			return
		}
	}

	return
}
