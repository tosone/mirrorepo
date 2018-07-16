package clone

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/satori/go.uuid"
	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/bash"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
	"gopkg.in/cheggaaa/pb.v2"
)

const serviceName = "clone"

var cloneLocker = new(sync.Mutex)

var currCloneID uint

var cloneList = map[uint]*models.Repo{}

// Initialize ..
func Initialize() {
	channel := make(chan taskmgr.ServiceCommand, 1)
	go func() {
		for control := range channel {
			switch control.Cmd {
			case "start":
				for _, repo := range cloneList {
					if control.TaskContent.(taskmgr.TaskContentClone).Repo.ID == repo.ID {
						return
					}
				}
				cloneList[control.TaskContent.(taskmgr.TaskContentClone).Repo.ID] = control.TaskContent.(taskmgr.TaskContentClone).Repo
				cloneLocker.Lock()
				clone(control.TaskContent.(taskmgr.TaskContentClone))
				delete(cloneList, control.TaskContent.(taskmgr.TaskContentClone).Repo.ID)
				cloneLocker.Unlock()
			case "stop":
				stop(control.TaskContent.(taskmgr.TaskContentClone).Repo.ID)
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

func clone(content taskmgr.TaskContentClone) {
	var err error
	var repo = content.Repo

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
			Time:   time.Now(),
		}
		if err = log.Create(); err != nil {
			logging.Error(err.Error())
		}
		if err = repo.UpdateByID(); err != nil {
			logging.Error(err.Error())
		}
	}()

	if com.IsDir(repo.RealPlace) {
		err = errors.New("dir is exist")
		return
	}

	currCloneID = repo.ID

	var address = repo.Address

	var cloneInfo = &bash.CloneInfo{
		Address:     address,
		Destination: repo.RealPlace,
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
	if err = detail(repo); err != nil {
		logging.Error(err)
	}
	if err = repo.UpdateByID(); err != nil {
		logging.Error(err)
	}
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
	if err = bash.RemoteReset(repo.RealPlace, repo.Address); err != nil {
		return
	}

	if repo.CommitCount, err = bash.CountCommits(repo.RealPlace); err != nil {
		return
	}

	var size uint64
	var commit string
	var historySize = new(models.HistorySize)
	var historyCommit = new(models.HistoryCommit)
	var uniqueID = uuid.NewV4().String()

	if size, err = bash.RepoSize(repo.RealPlace); err != nil {
		return
	}

	historySize.RepoID = repo.ID
	historySize.Size = size
	historySize.RepoLastID = uniqueID
	if err = historySize.Create(); err != nil {
		return
	}

	if commit, err = bash.CommitID(repo.RealPlace); err != nil {
		return
	}
	historyCommit.RepoID = repo.ID
	historyCommit.Commit = commit
	historyCommit.RepoLastID = uniqueID
	if err = historyCommit.Create(); err != nil {
		return
	}

	repo.Foreigner = uniqueID

	return
}
