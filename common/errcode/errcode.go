package errcode

import (
	"fmt"
)

var (
	// ErrNoSuchRecord ..
	ErrNoSuchRecord = fmt.Errorf("no such a record")
)
