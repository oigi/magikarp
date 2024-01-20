package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/oigi/Magikarp/api/v1"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("login", baseApi.Login) //用户登陆
	}
}
