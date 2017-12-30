package logging

import (
	"github.com/spf13/viper"
	"github.com/tosone/mirror-repo/models"
	"gopkg.in/lumberjack.v2"
)

var logger = new(lumberjack.Logger)
var sqlogger = new(lumberjack.Logger)

// Rotate rotate the output log file
func Rotate() {
	logger.Rotate()
}

var logLevel Level

func Setting() {
	sqlogger = &lumberjack.Logger{
		Filename:   viper.GetString("Log.Database"),
		MaxSize:    viper.GetInt("Log.MaxSize"),
		MaxBackups: viper.GetInt("Log.MaxBackups"),
		MaxAge:     viper.GetInt("Log.MaxAge"),
		LocalTime:  viper.GetBool("Log.LocalTime"),
		Compress:   viper.GetBool("Log.Compress"),
	}
	models.Logging(sqlogger)

	logger = &lumberjack.Logger{
		Filename:   viper.GetString("Log.App"),
		MaxSize:    viper.GetInt("Log.MaxSize"),
		MaxBackups: viper.GetInt("Log.MaxBackups"),
		MaxAge:     viper.GetInt("Log.MaxAge"),
		LocalTime:  viper.GetBool("Log.LocalTime"),
		Compress:   viper.GetBool("Log.Compress"),
	}

	logLevel = viper.Get("Log.Level").(Level)
}
