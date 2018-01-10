package models

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"github.com/tosone/mirror-repo/logging"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine
var err error

func Connect() {
	var dialString string
	if viper.GetString("Database.Engine") == "sqlite3" {
		dialString = viper.GetString("Database.Path")
	} else {
		dialString = fmt.Sprintf("%s://%s@%s:%s:%s",
			viper.GetString("Database.Engine"),
			viper.GetString("Database.Username"),
			viper.GetString("Database.Password"),
			viper.GetString("Database.Host"),
			viper.GetString("Database.Port"),
		)
	}

	engine, err = xorm.NewEngine(viper.GetString("Database.Engine"), dialString)
	if err != nil {
		logging.WithFields(logging.Fields{"engine": viper.GetString("Database.Engine")}).Panic(err.Error())
	}
	err = engine.Sync2(new(Repo), new(Log))
	if err != nil {
		logging.Panic(err.Error())
	}

	engine.ShowSQL(viper.GetBool("Log.ShowSQL"))
	engine.SetLogger(xorm.NewSimpleLogger(&lumberjack.Logger{
		Filename:   viper.GetString("Log.Database"),
		MaxSize:    viper.GetInt("Log.MaxSize"),
		MaxBackups: viper.GetInt("Log.MaxBackups"),
		MaxAge:     viper.GetInt("Log.MaxAge"),
		LocalTime:  viper.GetBool("Log.LocalTime"),
		Compress:   viper.GetBool("Log.Compress"),
	}))
}
