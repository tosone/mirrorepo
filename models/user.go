package models

import (
	"time"
)

type User struct {
	Id           int64  `xorm:"pk autoincr"`    // 主键
	Name         string `xorm:"notnull unique"` // 本地存储所用的名字
	Hash         string `xorm:"notnull"`        // 密码 hash
	Salt         string `xorm:"notnull"`        // 密码 salt
	SendMail     string `xorm:"notnull"`        // 是否发送邮件
	MailSMTP     string
	MailPort     int
	MailUserName string
	MailPassword string
	CreateAt     time.Time  `xorm:"created"` // 创建时间
	UpdateAt     time.Time  `xorm:"updated"` // 更新时间
	DeleteAt     *time.Time `xorm:"deleted"` // 删除时间
}

// Create ..
func (user *User) Create() (int64, error) {
	return engine.Insert(user)
}
