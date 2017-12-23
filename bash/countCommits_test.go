package bash

import (
	"fmt"
	"testing"
)

func Test_CountCommits(t *testing.T) {
	fmt.Println(CountCommits("/Users/tosone/awesome/bolt"))
	fmt.Println(FileCount("/Users/tosone/awesome/bolt"))
	Clone("")
}
