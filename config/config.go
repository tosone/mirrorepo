package config

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/Unknwon/com"
)

type config struct {
	DBEngine  string
	DBPath    string
	DBLog     bool
	MaxThread uint
	Repo      string
}

var Get = config{ // 默认配置
	DBEngine:  "sqlite3",
	DBPath:    "",
	DBLog:     true,
	MaxThread: 4,
	Repo:      "./repo",
}

func init() {
	if !com.IsDir(Get.Repo) {
		if err := os.MkdirAll(Get.Repo, 0755); err != nil {
			logrus.Error(err)
		}
	}
}
