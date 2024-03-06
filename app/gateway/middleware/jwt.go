package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/pkg/jwt"
    "github.com/oigi/Magikarp/pkg/resp"
    "net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
        // 这里假设Token放在Header的Authorization中，并使用Bearer开头
        // 这里的具体实现方式要依据你的实际业务情况决定
        authHeader := ctx.Request.Header.Get("Authorization")
        if authHeader == "" {
            ctx.JSON(http.StatusOK, gin.H{
                "code": 2003,
                "msg":  "请求头中auth为空",
            })
            ctx.Abort()
            return
        }

        // 解析 token
        claims, err := jwt.ParseToken(authHeader)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, resp.RespError(ctx, err, "未授权"))
            ctx.Abort()
            return
        }

        // 将当前请求的userID信息保存到请求的上下文c上
        ctx.Set("claims", claims)
        ctx.Next()
    }
}
