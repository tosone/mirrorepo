package errwebcode

import (
	"github.com/gin-gonic/gin"
)

var (
	// Normal ..
	Normal = gin.H{"code": 10000, "msg": ""}
	// AddressNull ..
	AddressNull = gin.H{"code": 10001, "msg": "address is null"}
	// DatabaseErr ..
	DatabaseErr = gin.H{"code": 10002, "msg": "database error"}
	// CloneCannotBeStopped ..
	CloneCannotBeStopped = gin.H{"code": 10003, "msg": "repo clone cannot be stopped"}
	// RepoIDNotValid ..
	RepoIDNotValid = gin.H{"code": 10004, "msg": "repo id is not valid"}
	// JSONMarshalErr ..
	JSONMarshalErr = gin.H{"code": 10005, "msg": ""}
	// DirExist ..
	DirExist = gin.H{"code": 10006, "msg": "the special dir is already exist"}
	// ServiceErr ..
	ServiceErr = gin.H{"code": 10007, "msg": "service error"}
)
