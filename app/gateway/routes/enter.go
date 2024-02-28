package routes

import (
    "github.com/oigi/Magikarp/app/gateway/routes/system"
    "github.com/oigi/Magikarp/app/gateway/routes/user"
)

type EnterGroup struct {
    System system.RouterGroup
    User   user.RouterGroup
}

var RouterGroupAPP = new(EnterGroup)
