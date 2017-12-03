package taskMgr

import (
	"github.com/Sirupsen/logrus"
	"github.com/tosone/mirror-repo/common/defination"
)

var taskList []defination.TaskChannel
var err error

// Register 注册一个服务
func Register(name string, channel chan defination.ServiceCommand) {
	taskList = append(taskList, defination.TaskChannel{Name: name, Channel: channel})
	err = Trigger(name)
	if err != nil {
		logrus.Error(err)
	}
}

// Trigger 触发一个服务
func Trigger(name string) error {
	return trans(name, defination.ServiceCommand{Cmd: "start"})
}

// Transport 传输信息到一个服务中
func Transport(name string, info defination.ServiceCommand) error {
	return trans(name, info)
}

func trans(name string, info defination.ServiceCommand) error {
	go func() {
		for _, t := range taskList {
			if t.Name == name {
				t.Channel <- info
				return
			}
		}
		logrus.Error("No such a service")
	}()
	return nil
}
