package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
)

func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq) (resp *favorite.FavoriteActionResp, err error) {
	return FavoriteClient.FavoriteAction(ctx, req)
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
	return FavoriteClient.FavoriteList(ctx, req)
}

func FavoriteCount(ctx context.Context, req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {
	return FavoriteClient.FavoriteCount(ctx, req)
}

func IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {
	return FavoriteClient.IsFavorite(ctx, req)
}
