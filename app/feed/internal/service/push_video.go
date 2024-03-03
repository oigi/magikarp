package service

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f *FeedServe) InitFollowFeed(ctx context.Context, req *feed.InitFollowFeedReq) (resp *feed.InitFollowFeedResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitFollowFeed not implemented")
}

func (f *FeedServe) PushRankVideo(ctx context.Context, req *feed.PushRankVideoReq) (resp *feed.PushRankVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRankVideo not implemented")
}

func (f *FeedServe) PushHotVideo(ctx context.Context, req *feed.PushHotVideoReq) (resp *feed.PushHotVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushHotVideo not implemented")
}

func (f *FeedServe) PushSimilarVideo(ctx context.Context, req *feed.PushSimilarVideoReq) (resp *feed.PushSimilarVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushSimilarVideo not implemented")
}

func (f *FeedServe) PushFollowVideo(ctx context.Context, req *feed.PushFollowVideoReq) (resp *feed.PushFollowVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushFollowVideo not implemented")
}

func (f *FeedServe) ListHistoryVideo(ctx context.Context, req *feed.ListHistoryVideoReq) (resp *feed.ListHistoryVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHistoryVideo not implemented")
}
