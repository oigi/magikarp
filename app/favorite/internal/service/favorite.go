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

func GetFeedServe() *FavoriteServe {
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
	}
	if len(cache) > 0 {
		resp.Code = e.SUCCESS
        resp.Msg = "查询成功"
        resp.VideoList =
        return
    }

}

func (f *FavoriteServe) IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {

}

func (f *FavoriteServe) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {

}

func convertVideoList(list []int64) (resp favorite.FavoriteListResp, err error) {
	var videoList []*feed.Video
	for _, i := range list {
		err := rpc.GetVideoById(context.Background(), &feed.QueryVideosReq{

		}, nil)
		if err != nil {
			continue
		}
		videoList = append(videoList, video)
	}
}