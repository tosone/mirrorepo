package models

import (
	"time"

	"github.com/tosone/mirror-repo/common/errCode"
	"github.com/tosone/mirror-repo/common/taskMgr"
)

type Task struct {
	Id       int64      `xorm:"pk autoincr"` // 主键
	RepoId   int64      `xorm:"index"`       // 外键
	Cmd      string     // 任务名
	Content  []byte     // 任务内容
	Progress int        // 进度
	Status   string     // 任务状态
	Error    string     // 失败原因
	CreateAt time.Time  `xorm:"created"` // 创建时间
	UpdateAt time.Time  `xorm:"updated"` // 更新时间
	DeleteAt *time.Time `xorm:"deleted"` // 删除时间
}

// Create ..
func (task *Task) Create() (int64, error) {
	defer func() {
		taskMgr.Trigger(task.Cmd)
	}()
	return engine.Insert(task)
}

// Failed ..
func (task *Task) Failed(err error) (int64, error) {
	var num int64
	task.Status = "failed"
	task.Error = err.Error()
	if num, err = task.Update(); err != nil || num == 0 {
		return num, err
	}
	return task.Delete()
}

// Success ..
func (task *Task) Success() (int64, error) {
	var num int64
	task.Status = "success"
	if num, err = task.Update(); err != nil || num == 0 {
		return num, err
	}
	return task.Delete()
}

// Delete ..
func (task *Task) Delete() (int64, error) {
	return engine.Delete(task)
}

func (task *Task) Update() (int64, error) {
	return engine.Id(task.Id).Update(task)
}

// Find ..
func (task *Task) Find() (r Repo, err error) {
	err = engine.Find(&r, task)
	return
}

func GetATask(cmd string) (*Task, error) {
	var b bool
	var err error
	var t = new(Task)
	b, err = engine.Where("cmd = ?", cmd).Limit(1).Get(t)
	if !b {
		err = errCode.ErrNoSuchRecord
	}
	return t, err
}
