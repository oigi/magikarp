package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/oigi/Magikarp/api/v1"
	"github.com/oigi/Magikarp/middleware"
)

type URouter struct{}

func (r *URouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.JWTAuthMiddleware())
	baseApi := v1.ApiGroupApp.UserApiGroup.BaseApi
	{
		userRouter.POST("login", baseApi.Login)       //用户登陆
		userRouter.POST("register", baseApi.Register) //用户注册
		userRouter.POST("")                           //用户改密
	}
}
