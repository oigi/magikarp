package user

import (
	"github.com/gin-gonic/gin"
)

type URouter struct{}

func (r *URouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use() //Todo

	{
		userRouter.POST("register", baseApi.Register) //用户注册
	}
}
