package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/oigi/Magikarp/api/v1"
)

type URouter struct{}

func (r *URouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.UserApiGroup.BaseApi
	{
		userRouter.POST("login", baseApi.Login) //用户登陆
	}
}
