package service

import (
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"sync"
)

var FeedServeOnce sync.Once
var FeedServeIns *FavoriteServe

type FavoriteServe struct {
	favorite.UnimplementedFavoriteServer
}

func GetFeedServe() *FavoriteServe {
	FeedServeOnce.Do(func() {
		FeedServeIns = &FavoriteServe{}
	})
	return FeedServeIns
}

//func (*FavoriteServe) SyncFavorite(ctx context.Context, req *favorite.SyncFavoriteReq) (resp *favorite.SyncFavoriteResp, err error) {
//	client := dao.NewFavoriteDao(ctx)
//	find := client.
//} TODO
