package models

import "github.com/jinzhu/gorm"

// HistoryCommit ..
type HistoryCommit struct {
	gorm.Model
	RepoID     uint
	RepoLastID string
	Commit     string
}

// Create ..
func (s *HistoryCommit) Create() error {
	return engine.Create(s).Error
}
