package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/models/logger"
)

var engine *gorm.DB

// Connect ..
func Connect() (err error) {
	var dialString string
	var engineType = viper.GetString("Database.Engine")
	if engineType == "sqlite3" {
		dialString = viper.GetString("Database.Path") + "?_busy_timeout=10000&_txlock=immediate"
	} else if engineType == "mysql" {
		dialString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			viper.GetString("Database.Username"),
			viper.GetString("Database.Password"),
			viper.GetString("Database.Host"),
			viper.GetString("Database.Port"),
			viper.GetString("Database.Database"),
		)
	} else if engineType == "postgres" {
		dialString = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
			engineType,
			viper.GetString("Database.Username"),
			viper.GetString("Database.Password"),
			viper.GetString("Database.Host"),
			viper.GetString("Database.Port"),
			viper.GetString("Database.Database"),
			viper.GetString("Database.SSLMode"),
		)
	} else {
		logging.Fatal(fmt.Sprintf("Not support this database: %s", engineType))
	}

	if engine, err = gorm.Open(viper.GetString("Database.Engine"), dialString); err != nil {
		logging.WithFields(logging.Fields{
			"engine":     viper.GetString("Database.Engine"),
			"dialString": dialString,
		}).Panic(err.Error())
	}

	if err = engine.AutoMigrate(
		new(Repo),
		new(Log),
		new(HistoryInfo),
	).Error; err != nil {
		logging.Panic(err.Error())
	}

	engine.LogMode(true)
	var gLogger logger.Logger
	engine.SetLogger(gLogger)
	return
}
