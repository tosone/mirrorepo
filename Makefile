BuildStamp = main.BuildStamp=$(shell date '+%Y-%m-%d_%I:%M:%S%p')
GitHash    = main.GitHash=$(shell git rev-parse HEAD)
Version    = main.Version=$(shell git describe --abbrev=0 --tags --always)
Target     = $(shell basename $(abspath $(dir $$PWD)))
Suffix     =

ifeq ($(OS),Windows_NT)
	OSName = windows
	Suffix = .exe
else
	OSName = $(shell uname -s)
endif

${OSName}: clean
	@GOOS=$(shell echo $@ | tr '[:upper:]' '[:lower:]') go build -v -o release/${Target}-$@${Suffix} -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}"

test: clean-test
	release/${Target}-${OSName}${Suffix} --config=config.yml scan $(GOPATH)/src/gopkg.in

test-web:
	release/${Target}-${OSName}${Suffix} --config=config.yml web

authors:
	echo "Authors\n=======\n\nProject's contributors:\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

clean:
	@$(RM) -r release

lint:
	gometalinter.v2 ./...

clean-test:
	$(RM) -r *.log mirror.db repo log
