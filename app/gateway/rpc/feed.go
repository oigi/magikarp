package rpc

//func ListFeed(ctx context.Context, req *feed.ListFeedReq) (resp *feed.ListFeedResp, err error) {
//	r, err := FeedClient.ListVideos(ctx, req)
//	if err != nil {
//		err = errors.WithMessage(err, "FeedClient.ListVideos error")
//		config.LOG.Error("", zap.Error(err))
//		return
//	}
//	if r.Code != e.SUCCESS {
//		err = errors.Wrap(errors.New("获取视频流失败"), "r.Code is unsuccessful")
//		config.LOG.Error("", zap.Error(err))
//		return
//	}
//	return r, err
//}
//
//func QueryVideos(ctx context.Context, req *feed.QueryVideosReq) (resp *feed.QueryVideosResp, err error) {
//	r, err := FeedClient.QueryVideos(ctx, req)
//	if err != nil {
//		err = errors.WithMessage(err, "FeedClient.QueryVideos error")
//		config.LOG.Error("", zap.Error(err))
//		return
//	}
//	if r.Code != e.SUCCESS {
//		err = errors.Wrap(errors.New("查询视频流失败"), "r.Code is unsuccessful")
//		config.LOG.Error("", zap.Error(err))
//		return
//	}
//	return r, err
//}
