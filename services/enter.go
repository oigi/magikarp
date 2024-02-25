package services

import (
	"github.com/oigi/Magikarp/services/common"
	"github.com/oigi/Magikarp/services/system"
	"github.com/oigi/Magikarp/services/user"
)

type EnterServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   user.ServiceGroup
	CommonServiceGroup common.ServiceGroup
}

var ServiceGroupApp = new(EnterServiceGroup)
