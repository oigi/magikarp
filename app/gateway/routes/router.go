package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/app/gateway/http"
    "github.com/oigi/Magikarp/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
    r := gin.Default()
    r.Use(middleware.Cors(), middleware.ErrorMiddleware())
    user := r.Group("/user")
    {
        user.POST("/login", http.UserLogin)
        user.POST("/register", http.UserRegister)
    }
    feed := r.Group("/feed").Use(middleware.JWTAuthMiddleware())
    {
        feed.POST("")
    }
    return r
}
