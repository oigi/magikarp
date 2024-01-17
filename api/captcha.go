package api

import (
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// Captcha 生成验证码
func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
}
