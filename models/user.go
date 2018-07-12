package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name         string // 本地存储所用的名字
	Hash         string // 密码 hash
	Salt         string // 密码 salt
	SendMail     string // 是否发送邮件
	MailSMTP     string
	MailPort     int
	MailUserName string
	MailPassword string
}

// Create ..
func (user *User) Create() error {
	return engine.Create(user).Error
}
