package user

import "github.com/oigi/Magikarp/services"

type ApiGroup struct {
	BaseApi
}

var (
	jwtService  = services.ServiceGroupApp.SystemServiceGroup.JwtService
	userService = services.ServiceGroupApp.UserServiceGroup.UService
)
