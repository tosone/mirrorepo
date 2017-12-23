package models

import "time"

type RepoAuthor struct {
	Id          int64      `xorm:"pk autoincr"` // 主键
	RepoID      uint       `xorm:"index"`       // 外键
	Author      string     `xorm:"notnull"`     // 作者的名字
	Email       string     `xorm:"notnull"`     // 作者的邮箱
	CommitCount uint       `xorm:"notnull"`     // 提交数量
	Percent     float64    `xorm:"notnull"`     // 百分比
	CreateAt    time.Time  `xorm:"created"`     // 创建时间
	UpdateAt    time.Time  `xorm:"updated"`     // 更新时间
	DeleteAt    *time.Time `xorm:"deleted"`     // 删除时间
}

// Create ..
func (author *RepoAuthor) Create() (int64, error) {
	return engine.Insert(&author)
}

// Delete ..
func (author *RepoAuthor) Delete() (int64, error) {
	return engine.Delete(&author)
}

// Find ..
func (author *RepoAuthor) Find() (r RepoAuthor, err error) {
	err = engine.Find(&r, author)
	return
}
