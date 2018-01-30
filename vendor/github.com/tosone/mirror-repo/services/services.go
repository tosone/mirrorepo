package services

import (
	"github.com/tosone/Mirror-repo/services/clone"
	"github.com/tosone/Mirror-repo/services/update"
)

func Initialize() {
	clone.Initialize()
	update.Initialize()
}
