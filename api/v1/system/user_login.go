package system

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/models"
    "github.com/oigi/Magikarp/models/request"
    "github.com/oigi/Magikarp/models/response"
    "github.com/oigi/Magikarp/utils"
    "net/http"
)

// Login 登陆 Todo 日志加入
func (b *BaseApi) Login(c *gin.Context) {
    var l request.Login
    err := c.ShouldBindJSON(&l)
    //ip := c.ClientIP()
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    // 判断验证码是否开启
    // 黑名单
    //解析json数据
    //验证码检查
    //登陆逻辑
    { // Todo 添加判断条件
        u := &models.User{Username: l.Username, Email: l.Email, Password: l.Password}

        user, err := userService.Login(u)
        if err != nil {
            // 登陆失败验证码次数+1
            response.FailWithMessage("用户名不存在或者密码错误", c)
            return
        }
        if user.Enable != 1 {
            response.FailWithMessage("用户被禁止登录", c)
            return
        }
        b.TokenNext(c, *user)
        //生成和返回JWT Token
    }

}

func (b *BaseApi) TokenNext(c *gin.Context, user models.User) {
    //j := &utils.JWT{} 唯一签名
    key := "啦啦啦啦啦" // Todo
    token, err := utils.GenerateJWT(user.Username, key)
    if err != nil {
        response.FailWithMessage("获取token失败", c)
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
