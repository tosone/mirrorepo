package models

import "errors"

var (
	// ErrNoSuchKey ..
	ErrNoSuchKey = errors.New("no suck a key")
	// ErrKeyAlreadyExist ..
	ErrKeyAlreadyExist = errors.New("the key already exist")
	// ErrDatabaseNull ..
	ErrDatabaseNull = errors.New("database is null")
)
