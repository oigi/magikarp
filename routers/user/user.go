package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("user")
	//
	//{
	//	userRouter.POST("admin_register", baseApi.Login) // 管理员注册账号
	//}

}
