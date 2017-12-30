package clone

import (
	"context"
	"encoding/json"
	"path"
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/kataras/go-errors"
	"github.com/spf13/viper"
	"github.com/tosone/mirror-repo/bash"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/errCode"
	"github.com/tosone/mirror-repo/common/taskMgr"
	"github.com/tosone/mirror-repo/logging"
	"github.com/tosone/mirror-repo/models"
	"gopkg.in/cheggaaa/pb.v1"
)

const serviceName = "clone"

var threadsClone = map[int64]threadInfo{}

type threadInfo struct {
	status    string
	progress  int
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func Initialize() {
	channel := make(chan defination.ServiceCommand, 1)
	go func() {
		for {
			select {
			case control := <-channel:
				switch control.Cmd {
				case "start":
					if int(len(threadsClone)) < viper.GetInt("Setting.MaxThread") {
						clone()
					} else {
						logging.Info("Clone has got its max thread.")
					}
				case "stop":
					stop(control.Id)
				}
			}
		}
	}()
	taskMgr.Register(serviceName, channel)
}

func clone() {
	var content defination.TaskContentClone
	var err error
	var task = &models.Task{}
	var repo = &models.Repo{}
	var ctx, ctxCancel = context.WithCancel(context.Background())
	defer func() {
		if err != nil {
			if err == errCode.ErrNoSuchRecord {
				logging.WithFields(logging.Fields{"serviceName": serviceName}).Error("Cannot find a valid task.")
			} else {
				if _, err = task.Failed(err); err != nil {
					logging.Error(err.Error())
				}
			}
		} else {
			var num int64
			if num, err = task.Success(); err != nil || num == 0 {
				logging.WithFields(logging.Fields{"err": err, "affectRows": num}).Error("Task done but delete error.")
			}
		}
	}()

	if task, err = models.GetATask(serviceName); err != nil {
		logging.Error(err.Error())
		return
	}

	if err = json.Unmarshal(task.Content, &content); err != nil {
		logging.Error(err.Error())
		return
	}

	repo.Id = task.RepoId

	if repo, err = repo.Find(); err != nil {
		logging.Error(err.Error())
		return
	}

	if _, err = repo.Create(); err != nil {
		logging.Error(err.Error())
		return
	}

	if com.IsDir(path.Join(content.Destination)) {
		err = errors.New("dir is exist")
		return
	}

	threadsClone[task.RepoId] = threadInfo{
		status:    "init",
		progress:  0,
		ctx:       ctx,
		ctxCancel: ctxCancel,
	}
	defer delete(threadsClone, task.RepoId)

	var cloneInfo = bash.CloneInfo{
		Address:     content.Address,
		Destination: content.Destination,
	}
	done := cloneInfo.Start()

	bar := pb.StartNew(100)
	go func() {
		for !bar.IsFinished() {
			bar.Set(cloneInfo.Progress)
			bar.Prefix(cloneInfo.Status)
			threadsClone[task.RepoId] = threadInfo{status: cloneInfo.Status, progress: cloneInfo.Progress}
			time.Sleep(time.Millisecond * 500)
			repo.Status = defination.RepoStatus(cloneInfo.Status)
			task.Progress = cloneInfo.Progress
			task.Update()
			repo.Update()
		}
	}()

	var wg = new(sync.WaitGroup)
	var doneResult error

	wg.Add(1)
	go func() {
		defer wg.Done()
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
		bar.Prefix("Error")
		bar.Set(100)
		bar.Finish()
		repo.Status = defination.RepoStatus(err.Error())
		repo.Update()
		return
	}

	bar.Prefix("Success")
	bar.Set(100)
	bar.Finish()
	repo.Status = defination.Success
	repo.Update()
}

func stop(id string) {
	for k, v := range threadsClone {
		if string(k) == id {
			if v.ctxCancel != nil {
				v.ctxCancel()
			}
			delete(threadsClone, k)
		}
	}
}
