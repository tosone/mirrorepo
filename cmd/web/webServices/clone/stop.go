package clone

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/errCode"
	"github.com/tosone/mirror-repo/models"
)

// Stop 停止
func Stop(repoID string) interface{} {
	var errRet = errCode.Normal
	var repo = models.Repo{}

	if id, err := strconv.Atoi(repoID); err != nil {
		logrus.Error(err)
		errRet = errCode.RepoIDNotValid
	} else {
		repo.ID = uint(id)
		if stopAble, err := repo.StopAble(); err != nil {
			logrus.Error(err)
			errRet = errCode.CloneCannotBeStopped
		} else {
			if !stopAble {
				errRet = errCode.CloneCannotBeStopped
			}
		}
	}

	return defination.WebServiceReturn{
		Code:  errRet.Code,
		Error: errRet.Describe,
	}
}
