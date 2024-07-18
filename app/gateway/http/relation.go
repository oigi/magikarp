package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/relation"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"go.uber.org/zap"
	"net/http"
)

func RelationAction(ctx *gin.Context) {
	var req relation.ActionReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r, err := rpc.RelationAction(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, r)
		return
	}
	ctx.JSON(http.StatusOK, r)
}

func RelationFollowList(ctx *gin.Context) {
	var req relation.FollowListReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r, err := rpc.RelationFollowList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, r)
	}
	ctx.JSON(http.StatusOK, r)
}

func RelationFollowerList(ctx *gin.Context) {
	var req relation.FollowerListReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r, err := rpc.RelationFollowerList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, r)
	}
	ctx.JSON(http.StatusOK, r)
}

func RelationFriendList(ctx *gin.Context) {
	var req relation.FriendListReq
	resp := model.CommonResp{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r, err := rpc.RelationFriendList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, r)
		return
	}
	ctx.JSON(http.StatusOK, r)
}
