package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/pkg/errors"
)

func IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {
	resp, err = FavoriteClient.IsFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return resp, nil
}

func FavoriteCount(ctx context.Context, req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {
	resp, err = FavoriteClient.FavoriteCount(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return resp, nil
}
