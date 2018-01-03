package clone

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/tosone/mirror-repo/bash"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/taskMgr"
	"github.com/tosone/mirror-repo/logging"
	"github.com/tosone/mirror-repo/models"
	"gopkg.in/cheggaaa/pb.v1"
)

const serviceName = "clone"

var cloneLocker = new(sync.Mutex)

var currCloneId int64

var cloneList = map[int64]*models.Repo{}

func Initialize() {
	channel := make(chan taskMgr.ServiceCommand, 1)
	go func() {
		for {
			select {
			case control := <-channel:
				switch control.Cmd {
				case "start":
					for _, repo := range cloneList {
						if control.TaskContent.(taskMgr.TaskContentClone).Repo.Id == repo.Id {
							return
						}
					}
					cloneList[control.TaskContent.(taskMgr.TaskContentClone).Repo.Id] = control.TaskContent.(taskMgr.TaskContentClone).Repo
					cloneLocker.Lock()
					clone(control.TaskContent.(taskMgr.TaskContentClone))
					delete(cloneList, control.TaskContent.(taskMgr.TaskContentClone).Repo.Id)
					cloneLocker.Unlock()
				case "stop":
					stop(control.TaskContent.(taskMgr.TaskContentClone).Repo.Id)
				}
			}
		}
	}()
	taskMgr.Register(serviceName, channel)
}

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

func clone(content taskMgr.TaskContentClone) {
	var err error
	var repo = content.Repo

	ctx, ctxCancel = context.WithCancel(context.Background())

	defer func() {
		var status = "success"
		var msg = ""
		if err != nil {
			status = "error"
			msg = err.Error()
		}
		log := &models.Log{
			RepoId: repo.Id,
			Cmd:    serviceName,
			Status: status,
			Msg:    msg,
			Time:   time.Now(),
		}
		_, err = log.Create()
		if err != nil {
			logging.Error(err.Error())
		}
		_, err = repo.Update()
		if err != nil {
			logging.Error(err.Error())
		}
	}()

	if com.IsDir(repo.RealPlace) {
		err = errors.New("dir is exist")
		return
	}

	currCloneId = repo.Id

	var address = repo.Address
	if content.Scan != "" {
		address = content.Scan
	}
	var cloneInfo = &bash.CloneInfo{
		Address:     address,
		Destination: repo.RealPlace,
	}
	done := cloneInfo.Start()
	bar := pb.StartNew(100)
	defer bar.Finish()
	var wg = new(sync.WaitGroup)
	var signalDone = make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			bar.Set(cloneInfo.Progress)
			bar.Prefix(repo.Name + " " + cloneInfo.Status)
			time.Sleep(time.Millisecond * 500)
			repo.Status = defination.RepoStatus(cloneInfo.Status)
			repo.Update()
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
		logging.Error(doneResult.Error())
		bar.Set(100)
		bar.Prefix(repo.Name + " " + "Error")
		repo.Status = defination.Error
		return
	}

	bar.Set(100)
	bar.Prefix(repo.Name + " " + "Success")
	repo.Status = defination.Success
	detail(repo)
}

func stop(id int64) {
	if id != currCloneId {
		return
	}
	if ctxCancel != nil {
		ctxCancel()
	}
}

func detail(repo *models.Repo) {
	var err error
	repo.CommitCount, err = bash.CountCommits(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}
	repo.Size, err = bash.RepoSize(repo.RealPlace)
	if err != nil {
		logging.Error(err.Error())
	}
}
