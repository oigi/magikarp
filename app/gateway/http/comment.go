package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/pkg/resp"
	"go.uber.org/zap"
	"net/http"
)

func CommentCount(ctx *gin.Context) {
	var req comment.CommentCountReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	count, err := rpc.CommentCount(ctx, &req)
	if err != nil {
		config.LOG.Error("调用RPC服务错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "调用RPC服务错误"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, count))
}

func CommentList(ctx *gin.Context) {
	var req comment.CommentListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	list, err := rpc.CommentList(ctx, &req)
	if err != nil {
		config.LOG.Error("调用RPC服务错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "调用RPC服务错误"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, list))
}

func CommentAction(ctx *gin.Context) {
	var req comment.CommentActionReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	result, err := rpc.CommentAction(ctx, &req)
	if err != nil {
		config.LOG.Error("调用RPC服务错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "调用RPC服务错误"))
	}
	ctx.JSON(http.StatusOK, result)
}
