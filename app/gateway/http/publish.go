package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/pkg/resp"
	"go.uber.org/zap"
	"net/http"
)

func CreateVideo(ctx *gin.Context) {
	var req publish.CreateVideoRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.CreateVideo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, r)
}

func ListVideo(ctx *gin.Context) {
	var req publish.ListVideoRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.ListVideo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}
