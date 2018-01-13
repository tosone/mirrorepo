package models

import (
	"github.com/jinzhu/gorm"
)

type RepoAuthor struct {
	gorm.Model
	RepoID      uint    // 外键
	Author      string  // 作者的名字
	Email       string  // 作者的邮箱
	CommitCount uint    // 提交数量
	Percent     float64 // 百分比
}

//// Create ..
//func (author *RepoAuthor) Create() (int64, error) {
//	return engine.Insert(&author)
//}
//
//// Delete ..
//func (author *RepoAuthor) Delete() (int64, error) {
//	return engine.Delete(&author)
//}
//
//// Find ..
//func (author *RepoAuthor) Find() (r RepoAuthor, err error) {
//	err = engine.Find(&r, author)
//	return
//}
