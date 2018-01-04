BuildStamp = main.BuildStamp=$(shell date '+%Y-%m-%d_%I:%M:%S%p')
GitHash    = main.GitHash=$(shell git rev-parse HEAD)
Version    = main.Version=$(shell git describe --abbrev=0 --tags)
Target     = mirror-repo

UNAME_S    = $(shell uname -s)

GOOS       = linux
CC         = gcc
Subfix     = linux

ifeq ($(UNAME_S),Darwin)
	GOOS   = darwin
	CC     = clang
	Subfix = mac 
endif

build: clean
	mkdir release
	CGO_ENABLED=1 CC=$(CC) GOOS=$(GOOS) GOARCH=amd64 go build -v -o release/${Target}-$(Subfix) -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go

test: cleanTest
	./release/mirror-repo-mac --config=config.yaml scan /Users/tosone/gocode/src/gopkg.in

authors:
	echo "Authors\n=======\n\nProject's contributors:\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

clean:
	-rm -rf release

cleanTest:
	-rm -rf *.log mirror.db repo