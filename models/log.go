package models

import (
	"time"
)

type Log struct {
	Id       int64      `xorm:"pk autoincr"` // 主键
	RepoId   int64      `xorm:"index"`       // 外键
	Cmd      string     // 任务名称
	Status   string     // 是否成功
	Time     time.Time  // 事件产生时间
	CreateAt time.Time  `xorm:"created"` // 创建时间
	UpdateAt time.Time  `xorm:"updated"` // 更新时间
	DeleteAt *time.Time `xorm:"deleted"` // 删除时间
}
