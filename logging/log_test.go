package logging

import (
	"testing"
)

func TestMain(t *testing.T) {
	Set(Config{
		FileOutput:  true,
		Filename:    "./logs/log.log",
		MaxSize:     10,
		LoggerLevel: ErrorLevel,
		Compress:    true,
	})
	Info("info level")   // should not be output
	Debug("debug level") // should not be output
	Warn("warn level")   // should not be output
	WithFields(Fields{"name": "tosone", "address": "here"}).Error("test")
	WithFields(Fields{"name": "tosone", "address": "here"}).Error("test")
	WithFields(Fields{"name": "tosone", "address": "here"}).Error("test")
	WithFields(Fields{"name": "tosone", "address": "here"}).Error("test")
	WithFields(Fields{"name": "tosone", "address": "here"}).Error("test")
	WithFields(Fields{"name": "tosone", "address": "here"}).Error("test")
	//<-time.After(time.Second * 10)
	//Rotate()
	//WithFields(Fields{"name": "tosone", "address": "here"}).Error("test") // should be output
	//Fatal("fatal level")                                                  // should be output
	//Panic("panic level")                                                  // should be output
}
