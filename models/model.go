package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Model struct {
	ID       uuid.UUID  `redis:"id"`
	CreateAt time.Time  `redis:"create_at"`
	UpdateAt time.Time  `redis:"update_at"`
	DeleteAt *time.Time `redis:"delete_at"`
}
