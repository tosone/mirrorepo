package bash

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type CloneInfo struct {
	Address     string
	Status      string
	Progress    int
	Destination string
	cmd         *exec.Cmd
}

func (info *CloneInfo) Start() (channel chan error) {
	var err error

	channel = make(chan error, 1)

	if info.Address == "" || info.Destination == "" {
		err = errors.New("clone info is not correct")
	}

	cmd := exec.Command("git", "clone", "--bare", "--progress", info.Address, info.Destination)

	var stderrPipe io.ReadCloser

	cmd.Stderr = cmd.Stdout
	stderrPipe, err = cmd.StderrPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}

	go func() {
		var n int
		for {
			var b = make([]byte, 80)
			n, err = stderrPipe.Read(b)
			if err == io.EOF {
				err = nil
				return
			}
			if err != nil {
				return
			}
			if n == 0 {
				return
			}
			var reg *regexp.Regexp
			if reg, err = regexp.Compile(`Receiving\s+objects:\s+(\d+)%[\w\W]+`); err != nil {
				return
			} else {
				matches := reg.FindStringSubmatch(string(b))
				if len(matches) == 2 {
					if info.Status != "Resolving" {
						info.Status = "Receiving"
						var num int
						if num, err = strconv.Atoi(matches[1]); err == nil {
							if num > info.Progress {
								info.Progress = num
							}
						} else {
							return
						}
					}
				}
			}
			if reg, err := regexp.Compile(`Resolving\s+deltas:\s+(\d+)%[\w\W]+\)`); err != nil {
				return
			} else {
				matches := reg.FindStringSubmatch(string(b))
				if len(matches) == 2 {
					if info.Status == "Receiving" {
						info.Status = "Resolving"
						info.Progress = 0
					}
					var num int
					if num, err = strconv.Atoi(matches[1]); err == nil {
						if num > info.Progress {
							info.Progress = num
						}
					} else {
						return
					}
				}
			}
		}
	}()
	go func() {
		cmd.Wait()
		channel <- err
	}()
	return
}

func (info *CloneInfo) Stop() (err error) {
	if _, err = os.FindProcess(info.cmd.Process.Pid); err != nil {
		return
	} else {
		if err = info.cmd.Process.Kill(); err != nil {
			return
		}
	}
	return
}
