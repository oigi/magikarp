package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/relation"
	"github.com/oigi/Magikarp/pkg/resp"
	"go.uber.org/zap"
	"net/http"
)

func RelationAction(ctx *gin.Context) {
	var req relation.ActionReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.RelationAction(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, r)
}

func RelationFollowList(ctx *gin.Context) {
	var req relation.FollowListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.RelationFollowList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}

func RelationFollowerList(ctx *gin.Context) {
	var req relation.FollowerListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.RelationFollowerList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}

func RelationFriendList(ctx *gin.Context) {
	var req relation.FriendListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.RelationFriendList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess(ctx, r))
}
