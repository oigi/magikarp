package service

import (
    "context"
    "github.com/oigi/Magikarp/app/favorite/internal/dao"
    "github.com/oigi/Magikarp/app/gateway/rpc"
    "github.com/oigi/Magikarp/grpc/pb/favorite"
    "github.com/oigi/Magikarp/grpc/pb/feed"
    "github.com/oigi/Magikarp/pkg/consts/e"
    "sync"
)

var FeedServeOnce sync.Once
var FeedServeIns *FavoriteServe

type FavoriteServe struct {
    favorite.UnimplementedFavoriteServiceServer
}

func GetFavoriteServe() *FavoriteServe {
    FeedServeOnce.Do(func() {
        FeedServeIns = &FavoriteServe{}
    })
    return FeedServeIns
}

func (f *FavoriteServe) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq) (resp *favorite.FavoriteActionResp, err error) {
    if req.ActionType == 1 {
        err := FavoriteAction(ctx, req, true)
        if err != nil {
            resp.Code = e.ERROR
            resp.Msg = "点赞失败"
        } else if req.ActionType == 2 {
            err := FavoriteAction(ctx, req, false)
            if err != nil {
                resp.Code = e.ERROR
                resp.Msg = "取消点赞失败"
            } else {
                resp.Code = e.InvalidParams
                resp.Msg = "参数错误"
            }
        }
    }
    return
}

func (f *FavoriteServe) FavoriteList(ctx context.Context, req *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
    // 1.查缓存
    cache, err := dao.QueryFavoriteListInCache(ctx, req.UserId)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "查询缓存失败"
        return
    }
    if len(cache) > 0 {
        resp.Code = e.SUCCESS
        resp.Msg = "查询成功"
        videoList, err := convertVideoList(cache)
        if err != nil {
            return nil, err
        }

        resp.VideoList = videoList
        return resp, err
    }

    mongo, err := dao.NewMongoClient(ctx).QueryFavoriteListInMongo(req.UserId)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "查询MongoDB失败"
        return
    }
    videoList, err := convertVideoList(mongo)
    resp.Code = e.SUCCESS
    resp.Msg = "查询成功"
    resp.VideoList = videoList
    return
}

func (f *FavoriteServe) IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {
    // 查询用户点赞列表
    cache, err := dao.QueryFavoriteListInCache(ctx, req.UserId)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "查询缓存失败"
        return
    }

    if len(cache) > 0 {
        cacheMap := make(map[int64]bool)
        for _, videoID := range cache {
            cacheMap[videoID] = true
        }
        resp.Code = e.SUCCESS
        resp.Msg = "查询成功"
        resp.IsFavorite = cacheMap
        return
    }

    mongo, err := dao.NewMongoClient(ctx).QueryFavoriteListInMongo(req.UserId)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "查询MongoDB失败"
        return
    }

    mongoMap := make(map[int64]bool)
    for _, videoID := range mongo {
        mongoMap[videoID] = true
    }
    resp.Code = e.SUCCESS
    resp.Msg = "查询成功"
    resp.IsFavorite = mongoMap
    return
}

func (f *FavoriteServe) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {
    if len(req.VideoIdList) == 0 {
        resp.Code = e.ERROR
        resp.Msg = "错误"
        return
    }
    count, err := dao.NewFavoriteDao(ctx).UpdateFavoriteCount(req)
    if err != nil {
        return nil, err
    }
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "更新点赞数失败"
    }
    resp.Code = e.SUCCESS
    resp.Msg = "更新成功"
    resp.VideoFavoriteCount = count
    return
}

func convertVideoList(list []int64) (videoList []*feed.Video, err error) {
    for _, i := range list {
        video, err := rpc.GetVideoById(context.Background(), &feed.QueryVideosReq{
            VideoId:  i,
            SearchId: 0,
        })
        if err != nil {
            continue
        }
        videoList = append(videoList, video)
    }

    return videoList, nil
}
