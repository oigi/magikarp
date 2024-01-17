package services

import "github.com/oigi/Magikarp/services/user"

type EnterServiceGroup struct {
	UserServiceGroup user.UserService
}

var ServiceGroupApp = new(EnterServiceGroup)
