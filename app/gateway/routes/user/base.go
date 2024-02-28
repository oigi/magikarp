package user

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/app/gateway/middleware"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
    baseRouter := Router.Group("user214").Use(middleware.JWTAuthMiddleware())
    {
        baseRouter.POST("login", baseApi.Login) //用户注册
    }
}
