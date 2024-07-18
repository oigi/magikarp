package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/http"
	"github.com/oigi/Magikarp/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors(), middleware.ErrorMiddleware())
	douyin := r.Group("/douyin")
	{
		user := douyin.Group("/user")
		{
			user.POST("/login/", http.UserLogin)
			user.POST("/register/", http.UserRegister)
		}
		getUser := douyin.Group("/user").Use(middleware.JWTAuthMiddleware())
		{
			getUser.GET("/", http.GetUserInfo)
		}
		feed := douyin.Group("/feed")
		{
			feed.GET("/", http.ListFeed)
		}
		comment := douyin.Group("/comment").Use(middleware.JWTAuthMiddleware())
		{
			comment.POST("action/", http.CommentAction)
			comment.GET("count/", http.CommentCount)
			comment.GET("list/", http.CommentList)
		}

		favorite := douyin.Group("/favorite").Use(middleware.JWTAuthMiddleware())
		{
			favorite.POST("action/", http.FavoriteAction)
			favorite.GET("list/", http.FavoriteList)
			favorite.GET("is/", http.IsFavorite)
		}
		favcount := douyin.Group("/favorite")
		{
			favcount.GET("/count/", http.FavoriteCount)
		}
		publish := douyin.Group("/publish").Use(middleware.JWTAuthMiddleware())
		{
			publish.GET("list/", http.ListVideo)
		}
		publishs := douyin.Group("/publish")
		{
			publishs.POST("action/", http.CreateVideo)
		}
		relation := douyin.Group("/relation").Use(middleware.JWTAuthMiddleware())
		{
			relation.POST("action/", http.RelationAction)
			relation.GET("follow/", http.RelationFollowList)
			relation.GET("follower/", http.RelationFollowerList)
			relation.GET("friend/", http.RelationFriendList)
		}

	}

	return r
}
