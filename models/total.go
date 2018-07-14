package models

import "github.com/jinzhu/gorm"

// Total ..
type Total struct {
	gorm.Model
	RepoNum  int
	Size     uint64
	LastSize uint64
}
