package mail

import (
	"testing"
	"time"
)

func TestSendMail(t *testing.T) {
	var a = Info{
		Repo: "repo",
		Size: 11,
		Time: time.Now(),
	}
	a.SendMail()
}
