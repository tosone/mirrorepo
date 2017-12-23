package web

import (
	"github.com/gin-gonic/gin"
	servicesClone "github.com/tosone/mirror-repo/cmd/web/webServices/clone"
	"github.com/tosone/mirror-repo/models"
	"github.com/tosone/mirror-repo/services"
)

func init() {
	models.Connect()
}

func Initialize() {
	services.Initialize()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	clone := engine.Group("clone")
	{
		clone.POST("/start", servicesClone.Start)
		clone.GET("/stop/:id", servicesClone.Stop)
	}

	engine.Run()
}
