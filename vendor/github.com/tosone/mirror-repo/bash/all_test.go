package bash

import (
	"fmt"
	"testing"
)

func TestGetRemoteUrl(t *testing.T) {
	fmt.Println(GetRemoteUrl("/Users/tosone/awesome/bolt"))
}

func TestActiveDays(t *testing.T) {
	fmt.Println(ActiveDays("/Users/tosone/awesome/bolt"))
}

func TestCountCommits(t *testing.T) {
	fmt.Println(CountCommits("/Users/tosone/awesome/bolt"))
}

func TestIsRepo(t *testing.T) {
	fmt.Println(IsRepo("/Users/tosone/awesome/bolt"))
}

func TestFileCount(t *testing.T) {
	fmt.Println(FileCount("/Users/tosone/awesome/linux"))
}

func TestShortLog(t *testing.T) {
	fmt.Println(ShortLog("/Users/tosone/awesome/bolt"))
}

func TestCommitId(t *testing.T) {
	fmt.Println(CommitId("/Users/tosone/awesome/bolt"))
}
