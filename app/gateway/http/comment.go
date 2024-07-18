package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"go.uber.org/zap"
	"net/http"
)

func CommentCount(ctx *gin.Context) {
	var req comment.CommentCountReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	count, err := rpc.CommentCount(ctx, &req)
	if err != nil {
		config.LOG.Error("调用RPC服务错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, count)
	}
	ctx.JSON(http.StatusOK, count)
}

func CommentList(ctx *gin.Context) {
	var req comment.CommentListReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	list, err := rpc.CommentList(ctx, &req)
	if err != nil {
		config.LOG.Error("调用RPC服务错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, list)
	}
	ctx.JSON(http.StatusOK, list)
}

func CommentAction(ctx *gin.Context) {
	var req comment.CommentActionReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	result, err := rpc.CommentAction(ctx, &req)
	if err != nil {
		config.LOG.Error("调用RPC服务错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, result)
	}
	ctx.JSON(http.StatusOK, result)
}
