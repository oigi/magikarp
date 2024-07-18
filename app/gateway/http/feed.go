package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/oigi/Magikarp/pkg/jwt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func ListFeed(ctx *gin.Context) {
	var req feed.ListFeedReq
	var userId int64
	authHeader := ctx.Query("token")
	if authHeader == "" {
		userId = 0
	} else {
		claims, err := jwt.ParseToken(authHeader)
		if err != nil {
			config.LOG.Error("解析错误", zap.Error(err))
			resp := model.CommonResp{
				StatusCode: e.ERROR,
				StatusMsg:  "解析错误",
			}
			ctx.JSON(http.StatusUnauthorized, resp)
		}
		userId = claims.ID
	}
	req.UserId = userId
	latestTimeString := ctx.Query("latest_time")
	latestTimeInt, _ := strconv.ParseInt(latestTimeString, 10, 64)

	req.LastTime = latestTimeInt

	resp, err := rpc.ListFeed(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Code = e.DOUYINSUCCESS
	ctx.JSON(http.StatusOK, resp)
}
