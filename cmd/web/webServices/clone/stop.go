package clone

import (
	"github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"github.com/tosone/mirror-repo/common/defination"
	"github.com/tosone/mirror-repo/common/errCode"
	"github.com/tosone/mirror-repo/models"
)

// Stop 停止
func Stop(repoID string) interface{} {
	var errRet = errCode.Normal
	var repo = &models.Repo{}
	var repoUUID uuid.UUID
	var err error

	repoUUID, err = uuid.FromBytes([]byte(repoID))
	if err != nil {
		logrus.Error(err)
		return defination.WebServiceReturn{
			Code:  errCode.RepoIDNotValid.Code,
			Error: errCode.RepoIDNotValid.Describe,
		}
	}

	repo.ID = repoUUID
	repo, err = repo.GetByID()
	if err != nil {
		return defination.WebServiceReturn{Code: errRet.Code, Error: errRet.Describe}
	}

	if repo.Status != "receiving" {
		errRet = errCode.CloneCannotBeStopped
	}

	return defination.WebServiceReturn{
		Code:  errRet.Code,
		Error: errRet.Describe,
	}
}
