package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/Unknwon/com"
	"gopkg.in/yaml.v2"
)

var defaultConfig = config{ // 默认配置
	DBEngine:  "sqlite3",
	DBPath:    "",
	Log:       "",
	MaxThread: 4,
	Repo:      "./repo",
}

var err error
var defaultConfigPath = "/etc/mirror-repo/config.yaml"

func init() {
	reset(&defaultConfig)
}

func reset(conf *config) {
	DBEngine = conf.DBEngine
	DBPath = conf.DBPath
	DBLog = conf.Log
	MaxThread = conf.MaxThread
	Repo = conf.Repo
}

func ensure(conf *config) {
	if !com.IsDir(conf.Repo) {
		if err := os.MkdirAll(conf.Repo, 0755); err != nil {
			logrus.Error(err)
		}
	}
}

func Setting(file string) {
	if file == "" {
		file = defaultConfigPath
	}

	if !com.IsFile(file) {
		log.Fatalf("Cannot find the %s.", file)
	}

	var conf = new(config)
	var content []byte
	if content, err = ioutil.ReadFile(""); err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(content, conf)
	reset(conf)
	ensure(conf)
}
