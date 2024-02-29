package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/app/gateway/http"
    "github.com/oigi/Magikarp/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
    r := gin.Default()
    r.Use(middleware.Cors(), middleware.ErrorMiddleware())
    v1 := r.Group("/")
    {
        v1.POST("/user/login", http.UserLogin)
        v1.POST("/user/register", http.UserRegister)

    }

    return r
}
