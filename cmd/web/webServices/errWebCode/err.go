package errWebCode

import (
	"github.com/gin-gonic/gin"
)

var (
	Normal               = gin.H{"code": 10000, "msg": ""}
	AddressNull          = gin.H{"code": 10001, "msg": "address is null"}
	DatabaseErr          = gin.H{"code": 10002, "msg": "database error"}
	CloneCannotBeStopped = gin.H{"code": 10003, "msg": "repo clone cannot be stopped"}
	RepoIDNotValid       = gin.H{"code": 10004, "msg": "repo id is not valid"}
	JSONMarshalErr       = gin.H{"code": 10005, "msg": ""}
	DirExist             = gin.H{"code": 10006, "msg": "the special dir is already exist"}
)
