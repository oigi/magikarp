package service

import (
    "context"
    "github.com/oigi/Magikarp/app/feed/internal/dao"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/grpc/pb/feed"
    "github.com/oigi/Magikarp/pkg/consts/e"
    "go.uber.org/zap"
)

func (f *FeedServe) CreateVideo(ctx context.Context, req *feed.CreateVideoReq) (resp *feed.CreateVideoResp, err error) {
    resp.Code = e.SUCCESS
    err = dao.NewFeedDao(ctx).InsertVideo(req)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "insert video error"
        config.LOG.Error("insert video error", zap.Error(err))
    }
    return
}

func (f *FeedServe) DeleteVideo(ctx context.Context, req *feed.DeleteVideoReq) (resp *feed.DeleteVideoResp, err error) {
    resp.Code = e.SUCCESS
    err = dao.NewFeedDao(ctx).DeleteVideoById(req)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "delete video error"
        config.LOG.Error("delete video error", zap.Error(err))
    }
    return
}

func (f *FeedServe) SearchVideo(ctx context.Context, req *feed.SearchVideoReq) (resp *feed.SearchVideoResp, err error) {
    resp.Code = e.SUCCESS
    r, err := dao.NewFeedDao(ctx).g
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "insert video error"
        config.LOG.Error("insert video error", zap.Error(err))
    }
    resp = & {
        resp.Video = r.
        resp.Code = e.ERROR
        resp.Msg = "success"
    }
    return
}

