package clone

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tosone/mirror-repo/cmd/web/webServices/errWebCode"
	"github.com/tosone/mirror-repo/models"
)

// Stop 停止
func Stop(context *gin.Context) {
	context.JSON(200, errWebCode.Normal)

	var repo = &models.Repo{}
	var err error

	repo.Id, err = strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
		context.JSON(200, errWebCode.RepoIDNotValid)
		return
	}

	repo, err = repo.Find()
	if err != nil {
		log.Fatal(err)
		context.JSON(200, errWebCode.DatabaseErr)
		return
	}

	if repo.Status != "receiving" {
		context.JSON(200, errWebCode.CloneCannotBeStopped)
	}

	return
}
