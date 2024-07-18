package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
)

func CommentCount(ctx context.Context, req *comment.CommentCountReq) (resp *comment.CommentCountResp, err error) {
	return CommentClient.CommentCount(ctx, req)
}

func FavoriteCount(ctx context.Context, req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {
	return FavoriteClient.FavoriteCount(ctx, req)
}

func IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {
	return FavoriteClient.IsFavorite(ctx, req)
}
