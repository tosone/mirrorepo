package main

import (
	"log"
	"runtime"

	"github.com/tosone/mirror-repo/cmd"
	"github.com/tosone/mirror-repo/cmd/version"
)

// Version version
var Version = "no provided"

// BuildStamp BuildStamp
var BuildStamp = "no provided"

// GitHash GitHash
var GitHash = "no provided"

func init() {
	//logging.Setting()
}

func main() {
	if runtime.GOOS == "windows" {
		log.Fatalln("Mirror-repo not support windows just linux.")
	}

	version.Setting(Version, BuildStamp, GitHash)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
