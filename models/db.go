package models

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/tosone/mirror-repo/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Major *gorm.DB

func Initialize() {
	var err error
	if Major, err = gorm.Open(config.Get.DBEngine, config.Get.DBPath); err != nil {
		logrus.Error(err)
	}
	if err = Major.LogMode(config.Get.DBLog).Error; err != nil {
		logrus.Error(err)
	}
	//Major.SetLogger(gormrus.New())
	//var task []models.Task
	//var repo models.Repo
	//Major.Find(&repo)
	//Major.Model(&repo).Where(models.Task{Command: "command"}).Related(&task)
	//log.Println(task)
	if err = Major.AutoMigrate(new(Repo), new(Task), new(RepoAuthor)).Error; err != nil {
		logrus.Error(err)
	}
}
