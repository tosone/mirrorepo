package clone

import (
	"net/http"
	"path"

	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/cmd/web/webServices/errwebcode"
	"github.com/tosone/mirrorepo/common/defination"
	"github.com/tosone/mirrorepo/common/taskmgr"
	"github.com/tosone/mirrorepo/models"
)

// Start ..
func Start(context *gin.Context) {
	var err error
	var address = context.PostForm("address")
	var name = context.PostForm("name")
	//var isSendEmail = context.PostForm("isSendEmail")

	if address == "" {
		context.JSON(http.StatusOK, errwebcode.AddressNull)
		return
	}

	var repo = &models.Repo{
		Address:   address,
		Status:    defination.Waiting,
		Name:      name,
		AliasName: path.Join(viper.GetString("Setting.Repo"), name),
		Travel:    viper.GetInt("Setting.Travel"),
	}

	if err = repo.Create(); err != nil {
		logging.Error(err)
		context.JSON(http.StatusOK, errwebcode.DatabaseErr)
		return
	}

	if com.IsDir(path.Join(viper.GetString("Setting.Repo"), name)) {
		context.JSON(http.StatusOK, errwebcode.DirExist)
		return
	}

	err = taskmgr.Transport(taskmgr.ServiceCommand{
		Task:        "clone",
		Cmd:         "start",
		TaskContent: taskmgr.TaskContentClone{Repo: repo},
	})

	if err != nil {
		logging.Error(err.Error())
		context.JSON(200, errwebcode.ServiceErr)
		return
	}
}
