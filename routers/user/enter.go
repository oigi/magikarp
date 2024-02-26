package user

import v1 "github.com/oigi/Magikarp/api/v1"

type RouterGroup struct {
	BaseRouter
	URouter
}

var baseApi = v1.ApiGroupApp.UserApiGroup.BaseApi
