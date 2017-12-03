package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/tosone/mirror-repo/cmd/scan"
	"github.com/tosone/mirror-repo/cmd/web"
	"github.com/tosone/mirror-repo/config"
	"github.com/tosone/mirror-repo/models"
	"github.com/tosone/mirror-repo/services"
)

// Version version
var Version = "no provided"

// BuildStamp BuildStamp
var BuildStamp = "no provided"

// GitHash GitHash
var GitHash = "no provided"

func main() {

	if runtime.GOOS == "windows" {
		logrus.Panicln("Mirror-repo not support windows just linux.")
	}

	var configPath string
	flag.StringVar(&configPath, "c", "./config.json", "config's absolutely path")
	flag.Parse()

	models.Initialize()
	services.Initialize()

	var args = os.Args
	if len(args) == 1 {
		web.Initialize()
	} else {
		switch args[1] {
		case "web":
			web.Initialize()
		case "scan":
			if len(args) == 3 {
				scan.Initialize(args[2])
			} else {
				scan.Initialize("")
			}
		default:
			web.Initialize()
		}
	}

	config.Get.DBPath = "./mirror.db"

	//
	//var repo = models.Repo{
	//	Address: "https://github.com/jinzhu/gorm.git",
	//	Status:  "Waiting",
	//	Name:    "gorm",
	//}
	//models.Major.Create(&repo)
	//models.AddTask(repo, "clone", defination.TaskContentClone{Address: repo.Address, Destination: path.Join(config.Get.Repo, "gorm")})
}
