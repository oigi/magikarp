package initalize

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/middleware"
	"github.com/oigi/Magikarp/routers"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := routers.RouterGroupAPP.User

	Router.Use(middleware.Cors()) //放行全部

	{
		systemRouter.InitUserRouter(Router.Group("user"))
	}

	return Router
}
