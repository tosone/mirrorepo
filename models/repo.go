package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
	"github.com/tosone/mirror-repo/common/funcs"
)

type Repo struct {
	Model
	Address     string `redis:"address"`       // 仓库地址
	Name        string `redis:"name"`          // 本地存储所用的名字
	Schedule    uint   `redis:"schedule"`      // 仓库两次更新之间的时间间隔
	Status      string `redis:"status"`        // 仓库状态
	Progress    int    `redis:"progress"`      // 进度
	CommitCount uint   `redis:"commit_count"`  // commit 数量
	IsSendEmail bool   `redis:"is_send_email"` // 是否发送邮件
}

func (repo *Repo) GetByAddress() (exist bool, err error) {
	var keys []string
	var retRepo *Repo
	keys, err = redis.Strings(engine.Get().Do("keys", "*"))
	if err != nil {
		return
	}
	for _, key := range keys {
		if strings.HasPrefix(key, fmt.Sprintf("hm:%s", funcs.BucketName(repo))) {
			var values []interface{}
			values, err = redis.Values(engine.Get().Do("HGETALL", key))
			if err != nil {
				return
			}
			err = redis.ScanStruct(values, &retRepo)
			if err != nil {
				return
			}
			if retRepo.Address == repo.Address && retRepo.DeleteAt == nil {
				return
			}
		}
	}
	err = ErrNoSuchKey
	return
}

func (repo *Repo) Create() (err error) {
	_, err = repo.GetByID()
	if err == ErrNoSuchKey {
		var now = time.Now()
		repo.ID = uuid.NewV4()
		repo.CreateAt = now
		repo.UpdateAt = now
		_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(repo), repo.ID)).AddFlat(&repo)...)
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

func (repo *Repo) GetByID() (retRepo *Repo, err error) {
	var keys []string
	keys, err = redis.Strings(engine.Get().Do("keys", "*"))
	if err != nil {
		return
	}
	for _, key := range keys {
		if strings.HasPrefix(key, fmt.Sprintf("hm:%s", funcs.BucketName(repo))) {
			var values []interface{}
			values, err = redis.Values(engine.Get().Do("HGETALL", key))
			if err != nil {
				return
			}
			err = redis.ScanStruct(values, &retRepo)
			if err != nil {
				return
			}
			if retRepo.ID == repo.ID && retRepo.DeleteAt == nil {
				return
			}
		}
	}
	err = ErrNoSuchKey
	return
}

func (repo *Repo) RemoveByID() (err error) {
	_, err = repo.GetByID()
	if err != nil {
		return
	}
	now := time.Now()
	repo.DeleteAt = &now
	_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(repo), repo.ID)).AddFlat(&repo)...)
	if err != nil {
		return
	}
	return
}

func (repo *Repo) UpdateByID(newRepo Repo) (err error) {
	_, err = repo.GetByID()
	if err != nil {
		return
	}
	now := time.Now()
	repo.UpdateAt = now
	_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(repo), repo.ID)).AddFlat(&repo)...)
	if err != nil {
		return
	}
	return
}
