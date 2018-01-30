package errCode

import (
	"fmt"
)

var (
	ErrNoSuchRecord = fmt.Errorf("no such a record")
)
