package initalize

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/middleware"
	"github.com/oigi/Magikarp/routers"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	userRouter := routers.RouterGroupAPP.User

	Router.Use(middleware.Cors()) //放行全部

	{
		userRouter.InitUserRouter(Router.Group("/"))
	}

	return Router
}
