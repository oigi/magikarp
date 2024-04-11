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
	feed := r.Group("/feed")
	{
		feed.POST("", http.ListFeed)
	}
	comment := r.Group("/comment").Use(middleware.JWTAuthMiddleware())
	{
		comment.POST("", http.CommentAction)
		comment.GET("count", http.CommentCount)
		comment.GET("list", http.CommentList)
	}
	favorite := r.Group("/favorite").Use(middleware.JWTAuthMiddleware())
	{
		favorite.POST("", http.FavoriteAction)
		favorite.GET("count", http.FavoriteCount)
	}
	publish := r.Group("/publish").Use(middleware.JWTAuthMiddleware())
	{
		publish.POST("", http.CreateVideo)
		publish.GET("list", http.ListVideo)
	}
	relation := r.Group("/relation").Use(middleware.JWTAuthMiddleware())
	{
		relation.POST("", http.RelationAction)
		relation.GET("follow", http.RelationFollowList)
		relation.GET("follower", http.RelationFollowerList)
		relation.GET("friend", http.RelationFriendList)
	}

	return r
}
