package services

import "github.com/oigi/Magikarp/services/system"

type EnterServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(EnterServiceGroup)
