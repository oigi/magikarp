package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/comment"
)

func CommentCount(ctx context.Context, req *comment.CommentCountReq) (resp *comment.CommentCountResp, err error) {
	return CommentClient.CommentCount(ctx, req)
}

func CommentAction(ctx context.Context, req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	return CommentClient.CommentAction(ctx, req)
}

func CommentList(ctx context.Context, req *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	return CommentClient.CommentList(ctx, req)
}
