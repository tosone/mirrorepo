package logging

import (
	"io"
	"log"
	"os"

	"github.com/tosone/mirror-repo/models"
	"gopkg.in/lumberjack.v2"
	"github.com/spf13/viper"
)

func Setting() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logger := &lumberjack.Logger{
		Filename:   viper.GetString("log"),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	mw := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(mw)
	models.Logging(logger)
}
