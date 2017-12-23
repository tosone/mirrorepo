BuildStamp = main.BuildStamp=$(shell date '+%Y-%m-%d_%I:%M:%S%p')
GitHash    = main.GitHash=$(shell git rev-parse HEAD)
Version    = main.Version=$(shell git describe --abbrev=0 --tags)
Target     = mirror-repo

build: clean
	mkdir release
	CGO_ENABLED=1 CC=clang GOOS=darwin GOARCH=amd64 go build -v -o release/${Target}-mac -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go
	#CGO_ENABLED=1 CC=gcc-7 GOOS=linux GOARCH=amd64 go build -v -o release/${Target}-linux -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go
	#CXX=arm-none-eabi-g++ CC=arm-none-eabi-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm go build -v -o release/${Target}-linux -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go

test:
	./release/mirror-repo-mac --config=config.yaml scan 123

authors:
	echo "Authors\n=======\n\nProject's contributors:\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

clean:
	-rm -rf release