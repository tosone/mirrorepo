package clone

import (
	"path"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/errCode"
	"github.com/tosone/mirror-repo/config"
	"github.com/tosone/mirror-repo/models"
)

type Info struct {
	Address     string
	Name        string
	IsSendEmail string
}

func (info Info) Start() interface{} {
	var errRet = errCode.Normal
	if info.Address != "" {
		errRet = errCode.AddressNull
	} else {
		var repo = &models.Repo{Address: info.Address, Status: "Waiting", Name: info.Name}
		if err := repo.Create(); err != nil {
			logrus.Error(err)
			errRet = errCode.DatabaseErr
		} else {
			var taskContent []byte
			taskContent, err = json.Marshal(defination.TaskContentClone{
				Address:     repo.Address,
				Destination: path.Join(config.Get.Repo, info.Name),
			})
			if err != nil {
				return err
			}
			task := models.Task{
				RepoID:  repo.ID,
				Name:    "clone",
				Content: taskContent,
			}
			err = task.Create()
			if err != nil {
				return err
			}
		}
	}

	return defination.WebServiceReturn{
		Code:  errRet.Code,
		Error: errRet.Describe,
	}
}
