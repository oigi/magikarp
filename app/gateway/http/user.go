package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/pkg/jwt"
	"github.com/oigi/Magikarp/pkg/resp"
	"go.uber.org/zap"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var req user.UserRegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	register, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		config.LOG.Error("UserRegister RPC服务调用错误")
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "UserRegister RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, register)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var req user.UserLoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	login, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		config.LOG.Error("UserUserLogin RPC服务调用错误")
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "UserUserLogin RPC服务调用错误"))
		return
	}
	accessToken, refreshToken, err := jwt.GenerateJWT(login.User.Id, login.User.Email)
	if err != nil {
		config.LOG.Error("加密错误:", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "加密错误"))
		return
	}
	u := &model.TokenData{
		ID:           login.User.Id,
		Email:        login.User.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, u))
}
