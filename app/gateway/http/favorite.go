package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"go.uber.org/zap"
	"net/http"
)

func FavoriteAction(ctx *gin.Context) {
	var req favorite.FavoriteActionReq
	resp := model.CommonResp{}
	json := model.PublishAction{}
	err := ctx.ShouldBindQuery(&json)
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r, err := rpc.FavoriteAction(ctx, &req)
	if err != nil {
		resp := model.CommonResp{
			StatusCode: e.ERROR,
			StatusMsg:  "FavoriteAction RPC服务调用错误",
		}
		config.LOG.Error("FavoriteAction RPC服务调用错误", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r.StatusCode = e.DOUYINSUCCESS
	ctx.JSON(http.StatusOK, r)
}

func FavoriteList(ctx *gin.Context) {
	var req favorite.FavoriteListReq
	id := ctx.GetInt64("id")
	req.UserId = id

	r, err := rpc.FavoriteList(ctx, &req)
	if err != nil {

		resp := model.CommonResp{
			StatusCode: e.DOUYINSUCCESS,
			StatusMsg:  "", // TODO 临时使用
		}

		//config.LOG.Error("RPC服务调用错误", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
		return
	}
	r.StatusCode = e.DOUYINSUCCESS
	ctx.JSON(http.StatusOK, r)
}

func FavoriteCount(ctx *gin.Context) {
	var req favorite.FavoriteCountReq
	var list []int64
	list = append(list, 1)
	req.VideoIdList = list
	r, err := rpc.FavoriteCount(ctx, &req)
	if err != nil {
		//config.LOG.Error("RPC服务调用错误", zap.Error(err))
		ctx.JSON(http.StatusOK, r)
		return
	}
	ctx.JSON(http.StatusOK, r)
}

func IsFavorite(ctx *gin.Context) {
	var req favorite.IsFavoriteReq
	resp := model.CommonResp{}
	err := ctx.ShouldBindQuery(&req)
	id := ctx.GetInt64("id")
	if id == 0 {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "用户不存在"
	}
	req.UserId = id
	var VideoIdList []int64
	VideoIdList[0] = 1
	req.VideoIdList = VideoIdList
	if err != nil {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "绑定参数错误"
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
	}
	r, err := rpc.IsFavorite(ctx, &req)
	if err != nil {
		config.LOG.Error("RPC服务调用错误", zap.Error(err))
		ctx.JSON(http.StatusOK, r)
		return
	}
	ctx.JSON(http.StatusOK, r)
}
