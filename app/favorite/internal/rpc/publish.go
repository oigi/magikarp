package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/feed"
)

func GetVideoById(ctx context.Context, req *feed.QueryVideosReq) (resp *feed.Video, err error) {
	return FeedClient.GetVideoById(ctx, req)
}
