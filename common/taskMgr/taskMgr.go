package taskMgr

import (
	"fmt"
	"sync"

	"github.com/tosone/mirrorepo/models"
)

// TaskChannel 任务触发通道
type TaskChannel struct {
	Name    string
	Channel chan ServiceCommand
}

// ServiceCommand 各个服务之间的命令传递
type ServiceCommand struct {
	Task        string      // 任务名
	Cmd         string      // 任务命令
	TaskContent interface{} // 任务内容
}

// TaskContentClone 克隆任务
type TaskContentClone struct {
	Repo *models.Repo
}

// TaskContentUpdate 更新任务
type TaskContentUpdate struct {
	Repo *models.Repo
}

var taskList []TaskChannel

// Register 注册一个服务
func Register(name string, channel chan ServiceCommand) {
	taskList = append(taskList, TaskChannel{Name: name, Channel: channel})
}

// Transport 传输信息到一个服务中
func Transport(info ServiceCommand) (err error) {
	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for _, t := range taskList {
			if t.Name == info.Task {
				wg.Done()
				t.Channel <- info
				return
			}
		}
		err = fmt.Errorf("no Such a service: %s", info.Task)
		wg.Done()
	}()
	wg.Wait()
	return
}
