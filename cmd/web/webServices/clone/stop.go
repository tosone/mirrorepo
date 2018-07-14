package clone

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/cmd/web/webServices/errwebcode"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
)

// Stop 停止
func Stop(context *gin.Context) {
	context.JSON(200, errwebcode.Normal)

	var err error
	var repo = &models.Repo{}

	var repoID uint64
	repoID, err = strconv.ParseUint(context.Param("id"), 10, 0)
	repo.ID = uint(repoID)
	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errwebcode.RepoIDNotValid)
		return
	}

	repo, err = repo.Find()

	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errwebcode.DatabaseErr)
		return
	}

	err = taskmgr.Transport(taskmgr.ServiceCommand{
		Task:        "clone",
		Cmd:         "stop",
		TaskContent: taskmgr.TaskContentClone{Repo: repo},
	})

	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errwebcode.ServiceErr)
		return
	}

	if repo.Status != "receiving" {
		context.JSON(200, errwebcode.CloneCannotBeStopped)
	}
}
