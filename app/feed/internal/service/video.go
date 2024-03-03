package service

import (
	"context"
	"github.com/oigi/Magikarp/consts/e"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f *FeedServe) CreateVideo(ctx context.Context, req *feed.CreateVideoReq) (resp *feed.CreateVideoResp, err error) {
	resp = new(feed.CreateVideoResp)
	resp.Code = e.SUCCESS
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}

func (f *FeedServe) DeleteVideo(ctx context.Context, req *feed.DeleteVideoReq) (resp *feed.DeleteVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}

func (f *FeedServe) GetVideo(ctx context.Context, req *feed.GetVideoListReq) (resp *feed.GetVideoListResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}

func (f *FeedServe) SearchVideo(ctx context.Context, req *feed.SearchVideoReq) (resp *feed.SearchVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchVideo not implemented")
}

func (f *FeedServe) ShareVideo(ctx context.Context, req *feed.ShareVideoReq) (resp *feed.ShareVideoResp, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareVideo not implemented")
}
