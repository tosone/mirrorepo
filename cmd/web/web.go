package web

import (
	"github.com/gin-gonic/gin"
	servicesClone "github.com/tosone/mirror-repo/cmd/web/webServices/clone"
)

func Initialize() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	clone := engine.Group("clone")
	{
		clone.POST("/start", func(c *gin.Context) {
			c.JSON(200, servicesClone.Info{
				Address:     c.PostForm("address"),
				Name:        c.PostForm("name"),
				IsSendEmail: c.PostForm("isSendEmail"),
			}.Start())
		})
		clone.GET("/stop/:id", func(c *gin.Context) {
			c.JSON(200, servicesClone.Stop(c.Param("id")))
		})
	}

	//repo := engine.Group("repo")
	//{
	//	repo.GET("/list", func(c *gin.Context) {
	//
	//	})
	//	repo.GET("/:id", func(c *gin.Context) {
	//
	//	})
	//}
	//
	//sync := engine.Group("sync")
	//{
	//	sync.GET("/start/:id", func(c *gin.Context) {
	//
	//	})
	//	sync.GET("/stop/:id", func(c *gin.Context) {
	//
	//	})
	//	sync.GET("/status/:id", func(c *gin.Context) {
	//
	//	})
	//}

	engine.Run()
}
