package models

import "errors"

var (
	ErrNoSuchKey       = errors.New("no suck a key")
	ErrKeyAlreadyExist = errors.New("the key already exist")
	ErrDatabaseNull    = errors.New("database is null")
)
