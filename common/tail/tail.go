package tail

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"context"

	"github.com/Sirupsen/logrus"
)

type Info struct {
	Status    string
	Progress  int
	ctx       context.Context
	ctxCancel context.CancelFunc
	Filename  string
}

func (info *Info) Watch() {
	info.ctx, info.ctxCancel = context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-info.ctx.Done():
				return
			default:
				if out, err := exec.Command("cat", "-e", info.Filename).Output(); err != nil {
					logrus.Error(err)
				} else {
					info.Status = "Connecting"
					info.handle(string(out))
					<-time.After(time.Millisecond * 200)
				}
			}
		}
	}()
}

func (info *Info) Stop() {
	if info.ctxCancel != nil {
		info.ctxCancel()
	}
	info.ctxCancel = nil
}

func (info *Info) handle(str string) {
	for _, strDollar := range strings.Split(str, "$") {
		if strings.Contains(strDollar, "^M") {
			for _, strM := range strings.Split(str, "^M") {
				if reg, err := regexp.Compile(`Receiving\s+objects:\s+(\d+)%[\w\W]+`); err != nil {
					logrus.Error(err)
				} else {
					matches := reg.FindStringSubmatch(strM)
					if len(matches) == 2 {
						if info.Status != "Resolving" {
							info.Status = "Receiving"
							if num, err := strconv.Atoi(matches[1]); err == nil {
								if num > info.Progress {
									info.Progress = num
								}
							} else {
								logrus.Error(err)
							}
						}
					}
				}
				if reg, err := regexp.Compile(`Resolving\s+deltas:\s+(\d+)%[\w\W]+\)`); err != nil {
					logrus.Error(err)
				} else {
					matches := reg.FindStringSubmatch(strM)
					if len(matches) == 2 {
						if info.Status == "Receiving" {
							info.Status = "Resolving"
							info.Progress = 0
						}

						if num, err := strconv.Atoi(matches[1]); err == nil {
							if num > info.Progress {
								info.Progress = num
							}
						} else {
							logrus.Error(err)
						}
					}
				}
				logrus.Println(info)
			}
		}
	}
	//logrus.Println(info)
}
