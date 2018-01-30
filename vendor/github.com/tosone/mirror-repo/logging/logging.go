package logging

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger = new(lumberjack.Logger)

// Rotate rotate the output log file
func Rotate() {
	logger.Rotate()
}

var logLevel Level

// Setting set the logger output method
func Setting() {
	logger = &lumberjack.Logger{
		Filename:   viper.GetString("Log.App"),
		MaxSize:    viper.GetInt("Log.MaxSize"),
		MaxBackups: viper.GetInt("Log.MaxBackups"),
		MaxAge:     viper.GetInt("Log.MaxAge"),
		LocalTime:  viper.GetBool("Log.LocalTime"),
		Compress:   viper.GetBool("Log.Compress"),
	}

	logLevel = Level(viper.Get("Log.Level").(int))
}
