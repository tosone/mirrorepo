package models

import "github.com/jinzhu/gorm"

// HistorySize ..
type HistorySize struct {
	gorm.Model
	RepoID     uint
	RepoLastID string
	Size       uint64
}

// Create ..
func (s *HistorySize) Create() error {
	return engine.Create(s).Error
}
