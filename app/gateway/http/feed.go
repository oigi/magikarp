package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/pkg/resp"
	"go.uber.org/zap"
	"net/http"
)

// ListFeed feed流
func ListFeed(ctx *gin.Context) {
	var req feed.ListFeedReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	// TODO 补充逻辑
}
