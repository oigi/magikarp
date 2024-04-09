package service

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
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

}

func (f *FavoriteServe) FavoriteList(ctx context.Context, req *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {

}

func (f *FavoriteServe) IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {

}

func (f *FavoriteServe) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {

}
