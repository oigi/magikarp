package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/oigi/Magikarp/pkg/resp"
	"go.uber.org/zap"
	"net/http"
)

func FavoriteAction(ctx *gin.Context) {
	var req favorite.FavoriteActionReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.FavoriteAction(ctx, &req)
	if err != nil {
		config.LOG.Error("FavoriteAction RPC服务调用错误")
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "FavoriteAction RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, r)
}

func FavoriteList(ctx *gin.Context) {
	var req favorite.FavoriteListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.FavoriteList(ctx, &req)
	if err != nil {
		config.LOG.Error("RPC服务调用错误")
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用错误"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}

func FavoriteCount(ctx *gin.Context) {
	var req favorite.FavoriteCountReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.FavoriteCount(ctx, &req)
	if err != nil {
		config.LOG.Error("RPC服务调用错误")
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用错误"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}

func IsFavorite(ctx *gin.Context) {
	var req favorite.IsFavoriteReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.IsFavorite(ctx, &req)
	if err != nil {
		config.LOG.Error("RPC服务调用错误")
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用错误"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}
