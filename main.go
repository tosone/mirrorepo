package main

import (
	"log"
	"runtime"

	"github.com/tosone/mirrorepo/cmd"
	"github.com/tosone/mirrorepo/cmd/version"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Version version
var Version = "no provided"

// BuildStamp BuildStamp
var BuildStamp = "no provided"

// GitHash GitHash
var GitHash = "no provided"

func main() {
	if runtime.GOOS == "windows" {
		log.Fatalln("mirrorepo not support windows just linux.")
	}

	version.Setting(Version, BuildStamp, GitHash)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
