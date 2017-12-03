package clone

import (
	"path"

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
		var repo = models.Repo{Address: info.Address, Status: "Waiting", Name: info.Name}
		if err := repo.Add(); err != nil {
			logrus.Error(err)
			errRet = errCode.DatabaseErr
		} else {
			if err := models.AddTask(repo, "clone", defination.TaskContentClone{Address: repo.Address, Destination: path.Join(config.Get.Repo, info.Name)}); err != nil {
				logrus.Error(err)
				errRet = errCode.DatabaseErr
			}
		}
	}

	return defination.WebServiceReturn{
		Code:  errRet.Code,
		Error: errRet.Describe,
	}
}
