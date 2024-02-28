package user

import (
    "github.com/oigi/Magikarp/log/api/v1"
)

type RouterGroup struct {
    BaseRouter
    URouter
}

var baseApi = v1.ApiGroupApp.UserApiGroup.BaseApi
