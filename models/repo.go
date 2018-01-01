package models

import (
	"time"

	"github.com/tosone/mirror-repo/common/defination"
)

type Repo struct {
	Id           int64                 `xorm:"pk autoincr"`    // 主键
	Address      string                `xorm:"notnull unique"` // 仓库地址
	Name         string                `xorm:"notnull unique"` // 本地存储所用的名字
	RealPlace    string                `xorm:"notnull"`        // 真正存储的地方
	Travel       int                   `xorm:"notnull"`        // 仓库两次更新之间的时间间隔
	LastTraveled time.Time             // 仓库上次被更新的时间
	Status       defination.RepoStatus // 仓库状态
	CommitCount  uint64                // commit 数量
	Size         uint64                // 仓库大小
	SendEmail    bool                  // 是否发送邮件
	CreateAt     time.Time             `xorm:"created"` // 创建时间
	UpdateAt     time.Time             `xorm:"updated"` // 更新时间
	DeleteAt     *time.Time            `xorm:"deleted"` // 删除时间
}

// Create ..
func (repo *Repo) Create() (int64, error) {
	return engine.Insert(repo)
}

// Delete ..
func (repo *Repo) Delete() (int64, error) {
	return engine.Delete(&repo)
}

// Find ..
func (repo *Repo) Find() (*Repo, error) {
	var r = new(Repo)
	var err error
	_, err = engine.Id(repo.Id).Get(r)
	return r, err
}

func (repo *Repo) Update() (int64, error) {
	return engine.Id(repo.Id).Update(repo)
}

func (repo *Repo) GetAll() (repos []*Repo, err error) {
	_, err = engine.Get(repos)
	return
}
