package routers

import (
	"github.com/oigi/Magikarp/routers/system"
	"github.com/oigi/Magikarp/routers/user"
)

type EnterGroup struct {
	System system.RouterGroup
	User   user.RouterGroup
}

var RouterGroupAPP = new(EnterGroup)
