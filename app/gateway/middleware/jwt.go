package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/oigi/Magikarp/pkg/jwt"
	"go.uber.org/zap"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		//authHeader := ctx.Request.Header.Get("Authorization")
		authHeader := ctx.Query("token")
		//if authHeader == "" {
		//	authHeader = ctx.GetHeader("token") // 请求头获取
		//}
		//if authHeader == "" {
		//	authHeader = ctx.PostForm("token")
		//}
		//  请求体读取在各自方法里写
		if authHeader == "" {
			if authHeader == "" {
				resp := model.CommonResp{
					StatusCode: e.ErrorAuthNotFound,
					StatusMsg:  "未登录,token不存在",
				}
				ctx.JSON(http.StatusOK, resp)
				ctx.Abort()
				return
			}
		}

		fmt.Println("-----------------------------")
		fmt.Println(authHeader)

		// 解析 token
		claims, err := jwt.ParseToken(authHeader)
		if err != nil {
			config.LOG.Error("解析错误", zap.Error(err))
			resp := model.CommonResp{
				StatusCode: e.ERROR,
				StatusMsg:  "解析错误",
			}
			ctx.JSON(http.StatusUnauthorized, resp)
			ctx.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		ctx.Set("id", claims.ID)
		ctx.Next()
	}
}
