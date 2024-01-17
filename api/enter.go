package api

import "github.com/oigi/Magikarp/services"

type ApiGroup struct {
	UserApiGroup.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = services.ServiceGroupApp.UserServiceGroup
)
