BuildStamp = main.BuildStamp=`date '+%Y-%m-%d_%I:%M:%S%p'`
GitHash    = main.GitHash=`git rev-parse HEAD`
Version    = main.Version=`git describe --abbrev=0 --tags`
Target     = morror-repo

build: clean
	mkdir release
	GOOS="darwin" GOARCH="amd64" go build -v -o release/${Target}-mac -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go
	GOOS="linux" GOARCH="amd64" go build -v -o release/${Target}-linux -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go

test:
	go test .

authors:
	echo "Authors\n=======\n\nProject's contributors:\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

clean:
	-rm -rf release