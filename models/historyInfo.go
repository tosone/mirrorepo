package models

import "github.com/jinzhu/gorm"

// HistoryInfo ..
type HistoryInfo struct {
	gorm.Model
	RepoID     uint
	RepoLastID string
	Commit     string
	Size       uint64
}

// Create ..
func (s *HistoryInfo) Create() error {
	return engine.Create(s).Error
}

// Find ..
func (s *HistoryInfo) Find() (info *HistoryInfo, err error) {
	err = engine.Model(new(HistoryInfo)).Find(info, &HistoryInfo{RepoLastID: s.RepoLastID}).Error
	return
}
