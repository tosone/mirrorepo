package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
	"github.com/tosone/mirror-repo/common/funcs"
)

type Task struct {
	Model
	RepoID  uuid.UUID `json:"repo_id"` // 外键
	Name    string    `json:"name"`    // 任务名
	Content []byte    `json:"content"` // 任务内容
	Status  string    `json:"status"`  // 任务状态
	Error   string    `json:"error"`   // 失败原因
}

// GetATask 获取新的任务
func (task *Task) GetATask(name string) (retTask *Task, err error) {
	var keys []string
	keys, err = redis.Strings(engine.Get().Do("keys", "*"))
	if err != nil {
		return
	}
	for _, key := range keys {
		if strings.HasPrefix(key, fmt.Sprintf("hm:%s", funcs.BucketName(task))) {
			var values []interface{}
			values, err = redis.Values(engine.Get().Do("HGETALL", key))
			if err != nil {
				return
			}
			err = redis.ScanStruct(values, &retTask)
			if err != nil {
				return
			}
			if retTask.DeleteAt == nil {
				return
			}
		}
	}
	err = ErrDatabaseNull
	return
}

// AddTask 添加任务
func (task *Task) Create() (err error) {
	_, err = task.GetByID()
	if err == ErrNoSuchKey {
		var now = time.Now()
		task.ID = uuid.NewV4()
		task.CreateAt = now
		task.UpdateAt = now
		_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(task), task.ID)).AddFlat(&task)...)
		if err != nil {
			return
		}
		return
	}
	if err == nil {
		err = ErrKeyAlreadyExist
	}
	return
}

// Failed 任务失败
func (task *Task) Failed(errParams error) (err error) {
	if errParams == nil {
		task.Success()
	}
	task.Error = errParams.Error()
	task.Status = "failed"
	return task.UpdateByID()
}

// Success 任务成功
func (task *Task) Success() error {
	task.Status = "success"
	return task.UpdateByID()
}

// UpdateByID ..
func (task *Task) UpdateByID() (err error) {
	_, err = task.GetByID()
	if err != nil {
		return
	}
	now := time.Now()
	task.UpdateAt = now
	_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(task), task.ID)).AddFlat(&task)...)
	if err != nil {
		return
	}
	return
}

// GetByID ..
func (task *Task) GetByID() (retTask *Task, err error) {
	var keys []string
	keys, err = redis.Strings(engine.Get().Do("keys", "*"))
	if err != nil {
		return
	}
	for _, key := range keys {
		if strings.HasPrefix(key, fmt.Sprintf("hm:%s", funcs.BucketName(task))) {
			var values []interface{}
			values, err = redis.Values(engine.Get().Do("HGETALL", key))
			if err != nil {
				return
			}
			err = redis.ScanStruct(values, &retTask)
			if err != nil {
				return
			}
			if retTask.ID == task.ID && retTask.DeleteAt == nil {
				return
			}
		}
	}
	err = ErrNoSuchKey
	return
}

// Remove 删除任务
func (task *Task) RemoveByID() (err error) {
	_, err = task.GetByID()
	if err != nil {
		return
	}
	now := time.Now()
	task.DeleteAt = &now
	_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(task), task.ID)).AddFlat(&task)...)
	if err != nil {
		return
	}
	return
}
