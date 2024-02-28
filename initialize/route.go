package initialize

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/app/gateway/middleware"
    "github.com/oigi/Magikarp/app/gateway/routes"
)

func Routers() *gin.Engine {
    Router := gin.Default()
    userRouter := routes.RouterGroupAPP.User

    Router.Use(middleware.Cors()) //放行全部

    {
        userRouter.InitUserRouter(Router.Group("/"))
        userRouter.InitBaseRouter(Router.Group("/"))
    }

    return Router
}
