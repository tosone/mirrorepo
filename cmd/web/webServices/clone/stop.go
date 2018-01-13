package clone

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tosone/mirror-repo/cmd/web/webServices/errWebCode"
	"github.com/tosone/mirror-repo/common/taskMgr"
	"github.com/tosone/mirror-repo/logging"
	"github.com/tosone/mirror-repo/models"
)

// Stop 停止
func Stop(context *gin.Context) {
	context.JSON(200, errWebCode.Normal)

	var repo = &models.Repo{}
	var err error
	var repoID uint64
	repoID, err = strconv.ParseUint(context.Param("id"), 10, 0)
	repo.ID = uint(repoID)
	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errWebCode.RepoIDNotValid)
		return
	}

	repo, err = repo.Find()

	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errWebCode.DatabaseErr)
		return
	}

	err = taskMgr.Transport(taskMgr.ServiceCommand{
		Task:        "clone",
		Cmd:         "stop",
		TaskContent: taskMgr.TaskContentClone{Repo: repo},
	})

	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errWebCode.ServiceErr)
		return
	}

	if repo.Status != "receiving" {
		context.JSON(200, errWebCode.CloneCannotBeStopped)
	}

	return
}
