package http

/*// ListFeed feed流
func ListFeed(ctx *gin.Context) {
	var req feed.ListFeedReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		config.LOG.Error("绑定参数错误: ", zap.Error(err))
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "绑定参数错误"))
		return
	}
	// TODO 补充逻辑
	r, err := rpc.ListFeed(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespError(ctx, err, "RPC服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, r)
}

*/
