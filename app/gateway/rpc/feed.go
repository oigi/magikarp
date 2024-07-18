package rpc

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func GetVideoById(ctx context.Context, req *feed.QueryVideosReq) (resp *feed.Video, err error) {
	return FeedClient.GetVideoById(ctx, req)
}

func ListFeed(ctx context.Context, req *feed.ListFeedReq) (resp *feed.ListFeedResp, err error) {
	r, err := FeedClient.ListVideos(ctx, req)
	if err != nil {
		err = errors.WithMessage(err, "FeedClient.ListVideos error")
		config.LOG.Error("", zap.Error(err))
		return
	}
	if r.Code != e.SUCCESS {
		err = errors.Wrap(errors.New("获取视频流失败"), "r.Code is unsuccessful")
		config.LOG.Error("", zap.Error(err))
		return r, nil
	}
	return r, err
}
