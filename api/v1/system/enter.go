package system

import "github.com/oigi/Magikarp/services"

type ApiGroup struct {
	BaseApi
}

var (
	userService = services.ServiceGroupApp.SystemServiceGroup.UserService
)
