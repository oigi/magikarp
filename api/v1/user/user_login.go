package user

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/models/common/response"
	"github.com/oigi/Magikarp/models/system"
	"github.com/oigi/Magikarp/models/user"
	"github.com/oigi/Magikarp/models/user/request"
	response2 "github.com/oigi/Magikarp/models/user/response"
	"github.com/oigi/Magikarp/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func (b *BaseApi) Login(c *gin.Context) {
	var model request.Login
	err := c.ShouldBindJSON(model)
	key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(model, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证码判断
	openCaptcha := global.CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || store.Verify(model.CaptchaId, model.Captcha, true) {
		u := &user.User{Email: model.Email, Password: model.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
}

func (b *BaseApi) TokenNext(c *gin.Context, user user.User) {
	token, err := utils.GenerateJWT(user.Email, []byte(global.CONFIG.JWT.SigningKey))
	if err != nil {
		global.LOG.Error("获取token失败", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	if !global.CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response2.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: time.Now().Add(7*24*time.Hour).Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Email); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Email); err != nil {
			global.LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response2.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: time.Now().Add(7*24*time.Hour).Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Email); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response2.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: time.Now().Add(7*24*time.Hour).Unix() * 1000,
		}, "登录成功", c)
	}
}

// Register 用户注册账号
func (b *BaseApi) Register(c *gin.Context) {
	var r request.Register

	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &user.User{Email: r.Email, NickName: r.NickName, Password: r.Password, Enable: r.Enable}
	userReturn, err := userService.Register(user)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(response2.UserResponse{User: *userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(response2.UserResponse{User: *userReturn}, "注册成功", c)
}
