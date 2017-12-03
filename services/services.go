package services

import (
	"github.com/tosone/mirror-repo/services/clone"
	"github.com/tosone/mirror-repo/services/scan"
	"github.com/tosone/mirror-repo/services/update"
)

func Initialize() {
	clone.Initialize()
	scan.Initialize()
	update.Initialize()
}
