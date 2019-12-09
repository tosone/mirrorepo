package clone

import (
	"context"
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/unknwon/com"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"gopkg.in/cheggaaa/pb.v2"

	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/bash"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
)

const serviceName = "clone"

var currCloneID uint

var cloneList = map[uint]*models.Repo{}

// Initialize ..
func Initialize() {
	var cloneLocker = new(sync.Mutex)
	channel := make(chan taskmgr.ServiceCommand, 1)
	go func() {
		for control := range channel {
			var task = control.TaskContent.(taskmgr.TaskContentClone)
			switch control.Cmd {
			case "start":
				var uniqueTask = true
				for _, repo := range cloneList {
					if task.Repo.ID == repo.ID {
						uniqueTask = false
					}
				}
				if uniqueTask {
					cloneList[task.Repo.ID] = task.Repo
					cloneLocker.Lock()
					if err := clone(task); err != nil {
						logging.Error(err)
					}
					delete(cloneList, task.Repo.ID)
					cloneLocker.Unlock()
				}
			case "stop":
				stop(task.Repo.ID)
			}
		}
	}()
	taskmgr.Register(serviceName, channel)
}

// WaitAll wait all of the clone tasks are over
func WaitAll() {
	var done = make(chan bool)
	go func() {
		for {
			if len(cloneList) == 0 {
				done <- true
				break
			}
			time.Sleep(time.Second)
		}
	}()
	<-done
}

var ctx context.Context
var ctxCancel context.CancelFunc

func clone(content taskmgr.TaskContentClone) (err error) {
	var repo = content.Repo
	var realPlace = path.Join(viper.GetString("Setting.Repo"), repo.AliasName)

	ctx, ctxCancel = context.WithCancel(context.Background())
	defer func() {
		if ctxCancel != nil {
			ctxCancel()
		}
	}()

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
			return
		}
		if err = repo.UpdateByID(); err != nil {
			return
		}
	}()

	if com.IsDir(realPlace) {
		err = fmt.Errorf("dir is exist already: %s", realPlace)
		return
	}

	currCloneID = repo.ID

	var address = repo.Address

	var cloneInfo = &bash.CloneInfo{
		Address:     address,
		Destination: realPlace,
	}
	done := cloneInfo.Start()

	if !strings.HasPrefix(address, "git") && !strings.HasPrefix(address, "http") && !strings.HasPrefix(address, "ssh") && com.IsDir(address) {
		repo.Address, err = bash.GetRemoteURL(address)
		if err != nil {
			return
		}
	}

	bar := pb.StartNew(100)
	defer bar.Finish()

	var wg = new(sync.WaitGroup)
	var signalDone = make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			bar.Set("prefix", repo.Name+" "+cloneInfo.Status+" ")
			bar.SetCurrent(int64(cloneInfo.Progress))
			time.Sleep(time.Millisecond * 500)
			repo.Status = defination.RepoStatus(cloneInfo.Status)
			if err = repo.UpdateByID(); err != nil {
				logging.Error(err)
			}
			select {
			case <-signalDone:
				return
			default:
			}
		}
	}()

	var doneResult error

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			signalDone <- true
		}()
		for {
			select {
			case <-ctx.Done():
				if err = cloneInfo.Stop(); err != nil {
					logging.Error(err.Error())
				}
				return
			case doneResult = <-done:
				return
			}
		}
	}()
	wg.Wait()

	if doneResult != nil {
		logging.Error(doneResult)
		bar.Set("prefix", repo.Name+" "+"Error ")
		repo.Status = defination.Error
		return
	}

	bar.Set("prefix", repo.Name+" "+"Success ")
	repo.Status = defination.Success

	if err = bash.RemoteReset(realPlace, repo.Address); err != nil {
		return
	}
	if err = detail(repo); err != nil {
		return
	}
	if err = repo.UpdateByID(); err != nil {
		return
	}
	return
}

func stop(id uint) {
	if id != currCloneID {
		return
	}
	if ctxCancel != nil {
		ctxCancel()
	}
}

func detail(repo *models.Repo) (err error) {
	var realPlace = path.Join(viper.GetString("Setting.Repo"), repo.AliasName)

	if repo.CommitCount, err = bash.CountCommits(realPlace); err != nil {
		return
	}

	var size uint64
	var commit string
	var historyInfo = new(models.HistoryInfo)
	var uniqueID = uuid.NewV4().String()

	if size, err = bash.RepoSize(realPlace); err != nil {
		return
	}
	if commit, err = bash.CommitID(realPlace); err != nil {
		return
	}

	historyInfo.RepoID = repo.ID
	historyInfo.Size = size
	historyInfo.Commit = commit
	historyInfo.RepoLastID = uniqueID
	if err = historyInfo.Create(); err != nil {
		return
	}

	repo.HistoryInfoID = uniqueID

	return
}
