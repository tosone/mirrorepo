package clone

import (
	"net/http"
	"path"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tosone/mirrorepo/cmd/web/webServices/errWebCode"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskMgr"
	"github.com/tosone/mirrorepo/logging"
	"github.com/tosone/mirrorepo/models"
)

func Start(context *gin.Context) {
	var err error
	var address = context.PostForm("address")
	var name = context.PostForm("name")
	//var isSendEmail = context.PostForm("isSendEmail")

	if address == "" {
		context.JSON(http.StatusOK, errWebCode.AddressNull)
		return
	}

	var repo = &models.Repo{
		Address:   address,
		Status:    defination.Waiting,
		Name:      name,
		RealPlace: path.Join(viper.GetString("Setting.Repo"), name),
		Travel:    viper.GetInt("Setting.Travel"),
	}

	if err = repo.Create(); err != nil {
		logging.Error(err.Error())
		context.JSON(http.StatusOK, errWebCode.DatabaseErr)
		return
	}

	if com.IsDir(path.Join(viper.GetString("Setting.Repo"), name)) {
		context.JSON(http.StatusOK, errWebCode.DirExist)
		return
	}

	err = taskMgr.Transport(taskMgr.ServiceCommand{
		Task:        "clone",
		Cmd:         "start",
		TaskContent: taskMgr.TaskContentClone{Repo: repo},
	})

	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errWebCode.ServiceErr)
		return
	}

	return
}
