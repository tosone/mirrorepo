package clone

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
	"time"

	"github.com/Unknwon/com"
	"github.com/kataras/go-errors"
	"github.com/satori/go.uuid"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/errCode"
	"github.com/tosone/mirror-repo/common/tail"
	"github.com/tosone/mirror-repo/common/taskMgr"
	"github.com/tosone/mirror-repo/config"
	"github.com/tosone/mirror-repo/models"
	"gopkg.in/cheggaaa/pb.v1"
)

const serviceName = "clone"

var threadsClone = map[string]threadInfo{}

type threadInfo struct {
	Status   string
	Progress int
}

func Initialize() {
	uuid.NewV4().String()
	channel := make(chan defination.ServiceCommand, 1)
	var ctx context.Context
	var ctxCancel context.CancelFunc
	go func() {
		for {
			select {
			case control := <-channel:
				switch control.Cmd {
				case "start":
					if uint(len(threadsClone)) < config.MaxThread {
						ctx, ctxCancel = context.WithCancel(context.Background())
						clone(ctx)
					} else {
						log.Println("clone has got its max thread.")
					}
				case "stop":
					if ctxCancel != nil {
						ctxCancel()
					}
				}
			}
		}
	}()
	taskMgr.Register(serviceName, channel)
}

func clone(ctx context.Context) {
	var content defination.TaskContentClone
	var err error
	var task = &models.Task{}
	var repo = &models.Repo{}
	defer func() {
		if err != nil {
			if err == errCode.ErrNoSuchRecord {
				log.Printf("Cannot find a valid %s task.", serviceName)
			} else {
				if _, err := task.Failed(err); err != nil {
					log.Println(err)
				}
			}
		} else {
			if num, err := task.Success(); err != nil || num == 0 {
				log.Printf("Task done but delete with error: %s, effect rows: %d", err.Error(), num)
			}
		}
	}()

	if task, err = models.GetATask(serviceName); err != nil {
		log.Println(err)
		return
	}

	if err = json.Unmarshal(task.Content, &content); err != nil {
		log.Println(err)
		return
	}

	repo.Id = task.RepoId
	if repo, err = repo.Find(); err != nil {
		log.Println(err)
		return
	}

	if com.IsDir(path.Join(content.Destination)) {
		err = errors.New("dir is exist")
		return
	}

	var tmpFile = "/tmp/" + uuid.NewV4().String()

	if file, err := os.OpenFile(tmpFile, os.O_CREATE, 0644); err != nil {
		log.Println(err)
		return
	} else {
		file.Close()
	}

	defer func() {
		if err := os.Remove(tmpFile); err != nil {
			log.Println(err)
		}
	}()

	threadsClone[content.Address] = threadInfo{Status: "Init"}
	defer delete(threadsClone, content.Address)

	cmd := exec.Command("./bash/clone.sh", content.Address, content.Destination, tmpFile)
	var info = tail.Info{Filename: tmpFile}
	info.Watch()

	bar := pb.StartNew(100)
	go func() {
		for !bar.IsFinished() {
			bar.Set(info.Progress)
			bar.Prefix(info.Status)
			threadsClone[content.Address] = threadInfo{Status: info.Status, Progress: info.Progress}
			time.Sleep(time.Millisecond * 500)
			repo.Status = defination.RepoStatus(info.Status)
			task.Progress = info.Progress
			task.Update()
			repo.Update()
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Println(err)
	}

	var done = make(chan error)
	go func() {
		done <- cmd.Wait()
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-done:
				return
			}
		}
	}()
	wg.Wait()

	if _, err := os.FindProcess(cmd.Process.Pid); err != nil {
		if err := cmd.Process.Kill(); err != nil {
			log.Println(err)
		}
	}

	bar.Prefix("Success")
	bar.Set(100)
	bar.Finish()
	info.Stop()
	repo.Status = defination.Success
	repo.Update()
}
