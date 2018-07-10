package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	servicesClone "github.com/tosone/mirrorepo/cmd/web/webServices/clone"
	"github.com/tosone/mirrorepo/logging"
)

var err error

func Initialize() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello"})
	})

	clone := engine.Group("clone")
	{
		clone.POST("/start", servicesClone.Start)
		clone.GET("/stop/:id", servicesClone.Stop)
	}

	var listenAddress = fmt.Sprintf("%s:%s", viper.GetString("Web.Host"), viper.GetString("Web.Port"))

	logging.Info(fmt.Sprintf("Listening and serving HTTP on %s.", listenAddress))
	err = engine.Run(listenAddress)
	if err != nil {
		logging.Panic(err.Error())
	}
}
