package clone

import (
	"encoding/json"
	"log"
	"net/http"
	"path"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/tosone/mirror-repo/cmd/web/webServices/errWebCode"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/config"
	"github.com/tosone/mirror-repo/models"
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

	var repo = &models.Repo{Address: address, Status: defination.Waiting, Name: name}
	if _, err = repo.Create(); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, errWebCode.DatabaseErr)
		return
	}

	if com.IsDir(path.Join(config.Repo, name)) {
		context.JSON(http.StatusOK, errWebCode.DirExist)
		return
	}

	var taskContent []byte
	taskContent, err = json.Marshal(defination.TaskContentClone{
		Address:     repo.Address,
		Destination: path.Join(config.Repo, name),
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, errWebCode.JSONMarshalErr)
		return
	}

	task := &models.Task{
		RepoId:  repo.Id,
		Cmd:     "clone",
		Content: taskContent,
	}
	_, err = task.Create()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, errWebCode.DatabaseErr)
		return
	}
	context.JSON(http.StatusOK, errWebCode.Normal)
	return
}
