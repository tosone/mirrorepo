package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Repo    Repo   // 关联仓库
	RepoID  uint   // 外键
	Name    string // 任务名
	Content []byte // 任务内容
	Status  string // 任务状态
	Error   string // 失败原因
}

type RepoAuthor struct {
	gorm.Model
	Repo        Repo // 关联仓库
	RepoID      uint // 外键
	Author      string
	Email       string
	CommitCount uint
	Percent     float64
}
