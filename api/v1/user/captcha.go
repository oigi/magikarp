package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/models/common/response"
	response2 "github.com/oigi/Magikarp/models/user/response"
	"go.uber.org/zap"
	"time"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha 生成验证码
func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(response2.UserCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
