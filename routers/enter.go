package routers

import "github.com/oigi/Magikarp/routers/user"

type EnterGroup struct {
	User user.RouterGroup
}

var RouterGroupAPP = new(EnterGroup)
