package models

type RepoAuthor struct {
	Model
	RepoID      uint    `json:"created_at"` // 外键
	Author      string  `json:"created_at"` // 作者的名字
	Email       string  `json:"created_at"` // 作者的邮箱
	CommitCount uint    `json:"created_at"` // 提交数量
	Percent     float64 `json:"created_at"` // 百分比
}
