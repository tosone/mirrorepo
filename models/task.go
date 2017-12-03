package models

import (
	"encoding/json"

	"github.com/kataras/go-errors"
	"github.com/tosone/mirror-repo/common/taskMgr"
)

// GetATask 获取新的任务
func GetATask(name string) (task Task, err error) {
	err = Major.Not("status", []string{"Failed", "Success"}).Last(&task, Task{Name: name}).Error
	return
}

// AddTask 添加任务
func AddTask(repo Repo, name string, content interface{}) (err error) {
	var contentByte []byte
	if contentByte, err = json.Marshal(content); err != nil {
		return
	}
	var r Repo
	if err = Major.First(&r, repo).Error; err != nil {
		return
	}
	var task = Task{RepoID: r.ID, Name: name, Content: contentByte, Repo: repo}
	if err = Major.Create(&task).Error; err != nil {
		return
	}
	taskMgr.Trigger(name)
	return
}

// Failed 任务失败
func (task Task) Failed(err error) error {
	if err != nil {
		return errors.New("task not failed")
	}
	if err := Major.Where(task).Updates(Task{Status: "Failed", Error: err.Error()}).Error; err != nil {
		return err
	}
	return nil
}

// Success 任务成功
func (task Task) Success() error {
	if err := Major.Model(new(Task)).Where(task).Updates(Task{Status: "Success"}).Error; err != nil {
		return err
	}
	return nil
}

// Remove 删除任务
func (task Task) Remove() error {
	if err := Major.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}
