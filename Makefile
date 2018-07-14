BuildStamp = main.BuildStamp=$(shell date '+%Y-%m-%d_%I:%M:%S%p')
GitHash    = main.GitHash=$(shell git rev-parse HEAD)
Version    = $(shell git describe --abbrev=0 --tags --always)
Target     = $(shell basename $(abspath $(dir $$PWD)))
Suffix     =

ifeq ($(OS),Windows_NT)
	OSName = windows
	Suffix = .exe
else
	OSName = $(shell echo $(shell uname -s) | tr '[:upper:]' '[:lower:]')
endif

${OSName}: clean
	GOOS=$@ go build -v -o release/${Target}-$@${Suffix} -ldflags "-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

test: clean_test
	release/${Target}-${OSName}${Suffix} --config=config.yaml scan /Users/tosone/gocode/src/gopkg.in

authors:
	echo "Authors\n=======\n\nProject's contributors:\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

clean:
	$(RM) -r release

lint:
	gometalinter.v2 ./...

clean_test:
	$(RM) -r *.log mirror.db repo
