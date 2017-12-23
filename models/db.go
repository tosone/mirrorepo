package models

import (
	"io"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine
var err error

func Connect() {
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Panicln(err)
	}
	err = engine.Sync2(new(Repo), new(Task))
	if err != nil {
		log.Panicln(err)
	}
}

func Logging(logFile io.Writer) {
	engine.ShowSQL(true)
	engine.SetLogger(xorm.NewSimpleLogger(logFile))
}
