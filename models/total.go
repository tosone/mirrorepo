package models

import "time"

// Total ..
type Total struct {
	Id       int64 `xorm:"pk autoincr"` // 主键
	RepoNum  int
	Size     uint64
	LastSize uint64
	CreateAt time.Time  `xorm:"created"` // 创建时间
	UpdateAt time.Time  `xorm:"updated"` // 更新时间
	DeleteAt *time.Time `xorm:"deleted"` // 删除时间
}
