package routers

import "github.com/oigi/Magikarp/routers/system"

type EnterGroup struct {
	User system.RouterGroup
}

var RouterGroupAPP = new(EnterGroup)
