package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tosone/mirrorepo/common/defination"
)

// Repo ..
type Repo struct {
	gorm.Model
	Address         string                // 仓库地址
	Name            string                // 本地存储所用的名字
	RealPlace       string                // 真正存储的地方
	Travel          int                   // 仓库两次更新之间的时间间隔
	LastTraveled    time.Time             // 仓库上次被更新的时间
	Status          defination.RepoStatus // 仓库状态
	CommitId        string                // HEAD 的 CommitId
	LastCommitId    string                // 之前 HEAD 的 CommitId
	CommitCount     uint64                // commit 数量
	LastCommitCount uint64                // 之前 commit 数量
	Size            uint64                // 仓库大小
	LastSize        uint64                // 之前仓库大小
}

// Create ..
func (repo *Repo) Create() error {
	return engine.Create(repo).Error
}

// Delete ..
func (repo *Repo) Delete() error {
	return engine.Delete(&repo).Error
}

// Find ..
func (repo *Repo) Find() (r *Repo, err error) {
	err = engine.Where(repo.ID).Find(r).Error
	return
}

// UpdateByID ..
func (repo *Repo) UpdateByID() error {
	return engine.Model(new(Repo)).Where(repo.ID).Updates(repo).Error
}

// GetAll ..
func (repo *Repo) GetAll() (repos []*Repo, err error) {
	err = engine.Find(repos).Error
	return
}
