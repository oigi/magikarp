package v1

import (
	"github.com/oigi/Magikarp/api/v1/system"
	"github.com/oigi/Magikarp/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
