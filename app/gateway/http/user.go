package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/oigi/Magikarp/pkg/jwt"
	"go.uber.org/zap"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var req user.UserRegisterReq
	var json model.Login
	// 从查询参数中获取用户名和密码
	err := ctx.ShouldBindQuery(&json)
	if err != nil {
		resp := model.CommonResp{
			StatusCode: e.ERROR,
			StatusMsg:  "ShouldBind绑定错误",
		}
		ctx.JSON(http.StatusOK, resp)
	}
	req.Email = json.Username
	req.Password = json.Password

	register, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		config.LOG.Error("UserRegister RPC服务调用错误")
		ctx.JSON(http.StatusOK, register)
	}

	if register.StatusCode == e.SUCCESS {
		Token, err := jwt.GenerateJWT(register.UserId, req.Email)
		if err != nil {
			config.LOG.Error("加密错误:", zap.Error(err))
			register.StatusMsg = "加密错误"
			ctx.JSON(http.StatusOK, register)
		}
		register.StatusCode = e.DOUYINSUCCESS
		register.Token = Token
	}

	ctx.JSON(http.StatusOK, register)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var req user.UserLoginReq
	var json model.Login
	err := ctx.ShouldBindQuery(&json)
	if err != nil {
		resp := model.CommonResp{
			StatusCode: e.ERROR,
			StatusMsg:  "ShouldBind绑定错误",
		}
		ctx.JSON(http.StatusOK, resp)
	}
	req.Email = json.Username
	req.Password = json.Password

	login, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		config.LOG.Error("UserUserLogin RPC服务调用错误")
		ctx.JSON(http.StatusOK, login)
	}
	if login.StatusCode == e.SUCCESS {
		Token, err := jwt.GenerateJWT(login.UserId, req.Email)
		if err != nil {
			config.LOG.Error("加密错误:", zap.Error(err))
			login.StatusMsg = "加密错误"
			ctx.JSON(http.StatusOK, login)
		}
		login.StatusCode = e.DOUYINSUCCESS
		login.Token = Token
	}
	ctx.JSON(http.StatusOK, login)
}

func GetUserInfo(ctx *gin.Context) {
	var req user.GetUserByIdReq
	//token := ctx.Query("token")
	id := ctx.GetInt64("id")
	req.UserId = id
	resp, err := rpc.GetUserById(ctx, &req)
	if err != nil {
		config.LOG.Error("RPC GetUserById 调用错误", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	resp.StatusCode = e.DOUYINSUCCESS
	ctx.JSON(http.StatusOK, resp)
}
