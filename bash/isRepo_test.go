package bash

import (
	"fmt"
	"testing"
)

func Test_IsRepo(t *testing.T) {
	fmt.Println(IsRepo("/Users/tosone/awesome/bolt"))
}
