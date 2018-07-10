package services

import (
	"github.com/tosone/mirrorepo/services/clone"
	"github.com/tosone/mirrorepo/services/update"
)

// Initialize ..
func Initialize() {
	clone.Initialize()
	update.Initialize()
}
